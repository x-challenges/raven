package http

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// ErrorPresenter interface.
type ErrorPresenter interface {
	HTTPStatus() *echo.HTTPError
}

// NewErrorHandler check err and return serialized response if err implements HTTPError interface.
var NewErrorHandler = func(_ *Config) func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		var (
			code     int
			response interface{}
		)

		// custom error
		var he ErrorPresenter
		if errors.As(err, &he) {
			status := he.HTTPStatus()

			code = status.Code
			response = status.Message
		} else {
			// echo http error
			var he *echo.HTTPError
			if errors.As(err, &he) {
				code = he.Code
				response = map[string]interface{}{
					"message": he.Message,
				}
			} else {
				code = echo.ErrInternalServerError.Code

				response = map[string]interface{}{
					"message": echo.ErrInternalServerError.Message,
				}
			}
		}

		// log error to stdout
		if err != nil {
			c.Logger().Error(err)
		}

		// response
		if err := c.JSON(code, response); err != nil {
			c.Logger().Error(err)
		}
	}
}
