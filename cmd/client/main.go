package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func init() {
	var cfg zap.Config
	cfg = zap.NewDevelopmentConfig()

	// set time format
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	// add color for console
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger = zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zapcore.DebugLevel))
}

func preferDeltaTemporalitySelector(kind sdkmetric.InstrumentKind) metricdata.Temporality {
	switch kind {
	case sdkmetric.InstrumentKindCounter,
		sdkmetric.InstrumentKindObservableCounter,
		sdkmetric.InstrumentKindHistogram:
		return metricdata.DeltaTemporality
	default:
		return metricdata.CumulativeTemporality
	}
}

func initProvider(ctx context.Context) func() {
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceName("client"),
			semconv.ServiceNamespace("demo"),
			semconv.ServiceVersion("0.1.0"),
		),
	)
	if err != nil {
		logger.Fatal("failed to create resource", zap.Error(err))
	}

	otelAgentAddr, ok := os.LookupEnv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if !ok {
		otelAgentAddr = "127.0.0.1:4317"
	}
	metricExp, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(otelAgentAddr),
		otlpmetricgrpc.WithCompressor("gzip"),
		otlpmetricgrpc.WithTemporalitySelector(preferDeltaTemporalitySelector),
	)

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExp,
				sdkmetric.WithInterval(2*time.Second),
			),
		),
	)
	otel.SetMeterProvider(meterProvider)

	return func() {
		cxt, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		// pushes any last exports to the receiver
		if err := meterProvider.Shutdown(cxt); err != nil {
			otel.Handle(err)
		}
	}
}

var (
	rng       = rand.New(rand.NewSource(time.Now().UnixNano()))
	meter     metric.Meter
	appIDList = []string{"1024", "568", "106", "1025", "197"}
)

func main() {
	ctx := context.Background()
	otelShutdownHook := initProvider(ctx)
	defer otelShutdownHook()

	meter = otel.Meter("demo-server")

	requestLatency, _ := meter.Float64Histogram(
		"demo_client/request_latency",
		metric.WithDescription("The latency of requests processed"),
	)

	requestCount, _ := meter.Int64Counter(
		"demo_client/request_counts",
		metric.WithDescription("The number of requests processed"),
	)

	for {
		appID := appIDList[rng.Intn(len(appIDList))]
		startTime := time.Now()
		makeHelloRequest(ctx, appID)
		latencyMs := float64(time.Since(startTime)) / 1e6
		requestLatency.Record(ctx, latencyMs, metric.WithAttributes(attribute.String("interface", "hello"), attribute.String("app_id", appID)))
		requestCount.Add(ctx, 1, metric.WithAttributes(attribute.String("interface", "hello"), attribute.String("app_id", appID)))

		fmt.Printf("Latency: %.3fms\n", latencyMs)
		time.Sleep(time.Duration(1) * time.Second)

	}
}

func makeHelloRequest(ctx context.Context, appID string) {
	demoServerAddr, ok := os.LookupEnv("DEMO_SERVER_ENDPOINT")
	if !ok {
		demoServerAddr = "0.0.0.0:2333"
	}

	// Trace an HTTP client by wrapping the transport
	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	// Make sure we pass the context to the request to avoid broken traces.
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://%s/hello/%s", demoServerAddr, appID), nil)
	if err != nil {
		logger.Fatal("failed to create request", zap.Error(err))
	}

	// All requests made with this client will create spans.
	res, err := client.Do(req)
	if err != nil {
		logger.Fatal("failed to execute request", zap.Error(err))
	}
	res.Body.Close()
}
