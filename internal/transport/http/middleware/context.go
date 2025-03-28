package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luxixing/fx-gin/internal/domain"
	"github.com/rs/xid"
)

const (
	// RequestIdHeader is the header key for request ID
	RequestIdHeader = "X-Request-ID"
	// RealIPHeader is the header key for real IP address
	RealIPHeader = "X-Real-IP"
)

// RequestContext middleware adds request context information
func RequestContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request ID from header, if not exist, generate a new one
		requestID := c.GetHeader(RequestIdHeader)
		if requestID == "" || requestID == "undefined" {
			requestID = xid.New().String()
		}

		// Get real IP from header, if not exist, get from client IP
		realIP := c.GetHeader(RealIPHeader)
		if realIP == "" {
			realIP = c.ClientIP()
		}

		// Get content length from header, if not exist, get from request body size
		bodySize := c.Request.ContentLength
		contentLength := fmt.Sprintf("%.2fkb", float64(bodySize)/1024.0)

		// Create trace info
		trace := &domain.TraceInfo{
			RequestID:     requestID,
			ClientIP:      realIP,
			ContentLength: contentLength,
			StartTime:     time.Now(),
			Method:        c.Request.Method,
			Path:          c.Request.URL.Path,
		}

		// Set trace info to context
		c.Set(domain.TraceKey, trace)

		// Set response header
		c.Header(RequestIdHeader, requestID)

		c.Next()
	}
}
