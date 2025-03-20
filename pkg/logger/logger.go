package logger

import (
	"sync"
	"time"

	"github.com/luxixing/fx-gin/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Module = fx.Options(
		fx.Provide(NewLogger),
	)
	once   sync.Once
	logger *zap.Logger
)

func NewLogger(cfg *config.Config) *zap.Logger {
	once.Do(func() {
		var err error

		// Custom time encoder
		customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.00"))
		}

		// Custom encoder config
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = customTimeEncoder
		encoderConfig.StacktraceKey = "" // Disable stacktrace by default

		// Determine log level
		var level zapcore.Level
		switch cfg.Logger.Level {
		case "debug":
			level = zapcore.DebugLevel
		case "info":
			level = zapcore.InfoLevel
		case "warn":
			level = zapcore.WarnLevel
		case "error":
			level = zapcore.ErrorLevel
			encoderConfig.StacktraceKey = "stacktrace" // Enable stacktrace for error level
		default:
			level = zapcore.InfoLevel
		}

		// Create a production logger with custom encoder config and level
		logger, err = zap.Config{
			Encoding:         "json",
			EncoderConfig:    encoderConfig,
			Level:            zap.NewAtomicLevelAt(level),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}.Build()
		if err != nil {
			panic(err)
		}
	})
	return logger
}
