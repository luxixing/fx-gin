package utils

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

// LogEnv logs all environment variables using the provided logger
func LogEnv(logger *zap.Logger) {
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			logger.Info("ENV", zap.String(parts[0], parts[1]))
		}
	}
}
