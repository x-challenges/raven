package resty

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

// LoggerAdapter type.
type LoggerAdapter struct {
	logger *zap.Logger
}

// NewLoggerAdapter constructor.
func NewLoggerAdapter(logger *zap.Logger) resty.Logger {
	return &LoggerAdapter{
		logger: logger,
	}
}

// Debugf implements resty.Logger interface.
func (la *LoggerAdapter) Debugf(format string, v ...interface{}) {
	la.logger.Sugar().Debugf(format, v)
}

// Warnf implements resty.Logger interface.
func (la *LoggerAdapter) Warnf(format string, v ...interface{}) {
	la.logger.Sugar().Warnf(format, v)
}

// Errorf implements resty.Logger interface.
func (la *LoggerAdapter) Errorf(format string, v ...interface{}) {
	la.logger.Sugar().Errorf(format, v)
}
