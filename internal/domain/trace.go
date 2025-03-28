package domain

import "time"

const (
	// TraceKey is the key for trace information in context
	TraceKey = "trace"
)

// TraceInfo represents request trace information
type TraceInfo struct {
	RequestID     string    `json:"-"`              // Request ID
	ClientIP      string    `json:"client_ip"`      // Client IP address
	ContentLength string    `json:"content_length"` // Request body size
	StartTime     time.Time `json:"-"`              // Request start time
	Method        string    `json:"method"`         // HTTP method
	Path          string    `json:"path"`           // Request path
	//todo add user info
}
