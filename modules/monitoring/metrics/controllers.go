package metrics

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

// NewMetrics
func NewMetricsController(e *echo.Echo, config *Config) {
	e.GET(config.Monitoring.Metrics.ExporterURL, echoprometheus.NewHandler())
}
