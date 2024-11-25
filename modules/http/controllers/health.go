package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const HealthCheckPath = "/health"

func NewHealthCheckController(e *echo.Echo) {
	e.GET(HealthCheckPath, func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
