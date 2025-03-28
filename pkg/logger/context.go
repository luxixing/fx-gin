package logger

import (
	"context"

	"github.com/luxixing/fx-gin/pkg/utils"
	"go.uber.org/zap"
)

// WithTraceFields adds trace fields to the log
func WithTraceFields(ctx context.Context, fields ...zap.Field) []zap.Field {
	trace := utils.FromContext(ctx)
	if trace == nil {
		return fields
	}

	return append([]zap.Field{zap.String("request_id", trace.RequestID)}, fields...)
}

// Info logs an info message with trace fields
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = WithTraceFields(ctx, fields...)
	zap.L().With(fields...).Info(msg)
}

// Error logs an error message with trace fields
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = WithTraceFields(ctx, fields...)
	zap.L().With(fields...).Error(msg)
}

// Debug logs a debug message with trace fields
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fields = WithTraceFields(ctx, fields...)
	zap.L().With(fields...).Debug(msg)
}

// Warn logs a warning message with trace fields
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fields = WithTraceFields(ctx, fields...)
	zap.L().With(fields...).Warn(msg)
}

// Fatal logs a fatal message with trace fields
func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	fields = WithTraceFields(ctx, fields...)
	zap.L().With(fields...).Fatal(msg)
}
