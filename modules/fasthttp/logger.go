package fasthttp

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// LoggerAdapter
type LoggerAdapter struct {
	logger *zap.Logger
}

var _ fasthttp.Logger = (*LoggerAdapter)(nil)

func NewLoggerAdapter(logger *zap.Logger) *LoggerAdapter {
	return &LoggerAdapter{
		logger: logger,
	}
}

// Printf implements LoggerAdapter interface
func (la *LoggerAdapter) Printf(format string, args ...any) {
	la.logger.Sugar().Debugf(format, args...)
}
