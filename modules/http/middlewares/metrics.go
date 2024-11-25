package middlewares

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

type T string

type MetricsOptions struct {
	Skippers Skippers
}

type MetricsOption func(*MetricsOptions)

func WithMetricsSkippers(v ...Skipper) MetricsOption {
	return func(to *MetricsOptions) {
		to.Skippers = v
	}
}

var MetricsMiddleware = func(opts ...MetricsOption) func() echo.MiddlewareFunc {
	var (
		options = &MetricsOptions{}
	)

	for _, opt := range opts {
		opt(options)
	}

	return func() echo.MiddlewareFunc {
		return echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
			Skipper:   options.Skippers.Handler,
			Namespace: "http",
		})
	}
}
