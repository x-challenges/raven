package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type LoggerOptions struct {
	Skippers Skippers
}

type LoggerOption func(*LoggerOptions)

func WithLoggerSkippers(v ...Skipper) LoggerOption {
	return func(lo *LoggerOptions) {
		lo.Skippers = v
	}
}

var LoggerMiddleware = func(opts ...LoggerOption) func(logger *zap.Logger) echo.MiddlewareFunc {
	var (
		options = &LoggerOptions{}
	)

	for _, opt := range opts {
		opt(options)
	}

	return func(logger *zap.Logger) echo.MiddlewareFunc {
		return middleware.RequestLoggerWithConfig(
			middleware.RequestLoggerConfig{
				Skipper:   options.Skippers.Handler,
				LogURI:    true,
				LogStatus: true,
				LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
					logger.Debug("http_request",
						zap.String("uri", v.URI),
						zap.Int("status", v.Status),
					)
					return nil
				},
			},
		)
	}
}
