package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapio"

	"github.com/x-challenges/raven/modules/http/serializers"
	"github.com/x-challenges/raven/modules/http/validators"
)

type ServerMuxParams struct {
	fx.In

	Logger      *zap.Logger
	Config      *Config
	Middlewares []echo.MiddlewareFunc `group:"http:middlewares"`
}

func NewServerMux(p ServerMuxParams) *echo.Echo {
	var (
		e = echo.New()
	)

	e.Debug = p.Config.Debug

	e.HideBanner = true
	e.HidePort = true

	// setup zap logger as a echo logger backend
	e.Logger.SetOutput(
		&zapio.Writer{
			Log: p.Logger,
		},
	)

	// custom json serializer via jsoniter
	e.JSONSerializer = serializers.NewJSONSerialzer(p.Config.Debug)

	// custom validation
	e.Validator = validators.NewValidator()

	// custom error handler
	e.HTTPErrorHandler = NewErrorHandler(p.Config)

	e.Use(p.Middlewares...)

	e.IPExtractor = echo.ExtractIPFromXFFHeader()

	return e
}

func NewServer(config *Config, e *echo.Echo) *http.Server {
	server := &http.Server{
		Addr:              config.HTTP.Bind,
		Handler:           e,
		ReadHeaderTimeout: 5 * time.Second,
	}

	server.SetKeepAlivesEnabled(config.HTTP.KeepAlive)

	return server
}
