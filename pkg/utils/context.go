package utils

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/luxixing/fx-gin/internal/domain"
)

// WithContext creates a context with trace information from gin.Context
func WithContext(c *gin.Context) context.Context {
	trace, exists := c.Get(domain.TraceKey)
	if !exists {
		return context.Background()
	}

	traceInfo := trace.(*domain.TraceInfo)
	return context.WithValue(context.Background(), domain.TraceKey, traceInfo)
}

// FromContext retrieves trace information from context
func FromContext(ctx context.Context) *domain.TraceInfo {
	trace, ok := ctx.Value(domain.TraceKey).(*domain.TraceInfo)
	if !ok {
		return nil
	}
	return trace
}
