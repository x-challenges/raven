package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"
)

type Skipper func(ctx echo.Context) bool

type Skippers []Skipper

func (s Skippers) Handler(ctx echo.Context) bool {
	for _, ss := range s {
		if ss(ctx) {
			return true
		}
	}
	return false
}

var (
	// skip metrics url.
	SkipMetrics Skipper = func(ctx echo.Context) bool {
		return strings.HasPrefix(ctx.Request().RequestURI, "/metrics")
	}

	// skip health check url.
	SkipHealthCheck Skipper = func(ctx echo.Context) bool {
		return strings.HasPrefix(ctx.Request().RequestURI, "/health")
	}
)

var (
	DefaultSkippers = []Skipper{
		SkipMetrics,
		SkipHealthCheck,
	}
)
