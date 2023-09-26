package log

import (
	"os"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once   sync.Once
	logger *zap.Logger
)

func Logger() *zap.Logger {
	logLevel, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevel = "INFO"
	}

	once.Do(func() {
		var cfg zap.Config
		logLevel, ok := ZapLevelMapping()[strings.ToUpper(logLevel)]
		if !ok {
			logLevel = zapcore.InfoLevel
		}
		if logLevel == zap.DebugLevel {
			cfg = zap.NewDevelopmentConfig()
		} else {
			cfg = zap.NewProductionConfig()
		}
		// set time format
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
		// add color for console
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger = zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel))
	})
	return logger
}

func ZapLevelMapping() map[string]zapcore.Level {
	return map[string]zapcore.Level{
		"DEBUG":  zapcore.DebugLevel,
		"INFO":   zapcore.InfoLevel,
		"WARN":   zapcore.WarnLevel,
		"ERROR":  zapcore.ErrorLevel,
		"DPANIC": zapcore.DPanicLevel,
		"PANIC":  zapcore.PanicLevel,
		"FATAL":  zapcore.FatalLevel,
	}
}
