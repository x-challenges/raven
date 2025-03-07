package pprof

import (
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo/v4"
)

// NewDebugController
func NewDebugController(e *echo.Echo, config *Config) {
	if config.Monitoring.Pprof.Enabled {
		pprof.Register(e)
	}
}
