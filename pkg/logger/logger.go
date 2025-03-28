package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	var config zap.Config
	// Determine log level based on APP_ENV
	env := os.Getenv("APP_ENV")
	switch env {
	case "dev", "test":
		config = zap.NewDevelopmentConfig()
	case "prod":
		config = zap.NewProductionConfig()
	default:
		config = zap.NewDevelopmentConfig()
	}

	// Configure encoder
	config.Encoding = "json" // console
	//config.EncoderConfig.TimeKey = "ts"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder

	// Disable stack trace
	config.DisableStacktrace = true

	logger, _ := config.Build(zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
}
