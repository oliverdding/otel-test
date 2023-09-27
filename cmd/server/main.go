package main

import (
	"context"
	"os"
	"time"

	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
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
			semconv.ServiceName("server"),
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
	meter        metric.Meter
	requestCount metric.Int64Counter
)

func main() {
	ctx := context.Background()
	otelShutdownHook := initProvider(ctx)
	defer otelShutdownHook()

	meter = otel.Meter("demo-server")

	app := fiber.New(fiber.Config{
		AppName:           "demo-server",
		CaseSensitive:     true,
		DisableKeepalive:  true,
		EnablePrintRoutes: true,
	})
	app.Use(fiberzap.New(fiberzap.Config{
		Logger:   logger,
		SkipURIs: []string{"/favicon.ico"},
		Fields:   []string{"time", "method", "path", "status", "respHeader:X-Request-ID", "ip", "port", "ua", "latency"},
		Levels:   []zapcore.Level{zapcore.WarnLevel},
	}))

	requestCount, _ = meter.Int64Counter(
		"request_counts",
		metric.WithDescription("The number of requests received"),
	)

	app.Get("/hello/:app_id", helloHandler)

	app.Listen("0.0.0.0:2333")
}

func helloHandler(c *fiber.Ctx) error {
	appID := c.Params("app_id")

	logger.Info("hello", zap.String("app_id", appID))

	requestCount.Add(c.UserContext(), 1, metric.WithAttributes(attribute.String("app_id", appID)))

	return nil
}
