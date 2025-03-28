package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luxixing/fx-gin/internal/domain"
	"go.uber.org/zap"
)

// Logger middleware using zap logger
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		trace, exists := c.Get(domain.TraceKey)
		if !exists {
			zap.S().Error("trace info not found in context")
			return
		}

		traceInfo := trace.(*domain.TraceInfo)
		latency := time.Since(traceInfo.StartTime).Milliseconds()

		zap.S().Infow("request info",
			"request_id", traceInfo.RequestID,
			"status", c.Writer.Status(),
			"latency_ms", latency,
			domain.TraceKey, traceInfo,
		)
	}
}
