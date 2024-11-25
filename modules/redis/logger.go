package redis

import (
	"context"

	"go.uber.org/zap"
)

// LoggerAdapter implements internal.Logging interface for redis client.
type LoggerAdapter struct {
	logger *zap.Logger
}

// NewLoggerAdapter construct.
func NewLoggerAdapter(logger *zap.Logger) *LoggerAdapter {
	return &LoggerAdapter{
		logger: logger,
	}
}

// Printf implements internal.Logging interface.
func (l *LoggerAdapter) Printf(_ context.Context, format string, v ...interface{}) {
	l.logger.Sugar().Infof(format, v)
}
