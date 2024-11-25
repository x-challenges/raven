package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type DumpOptions struct {
	Skippers Skippers
}

type DumpOption func(*DumpOptions)

func WithDumpSkippers(v ...Skipper) DumpOption {
	return func(do *DumpOptions) {
		do.Skippers = v
	}
}

var DumpMiddleware = func(opts ...DumpOption) func(logger *zap.Logger) echo.MiddlewareFunc {
	var options = &DumpOptions{}

	for _, opt := range opts {
		opt(options)
	}

	return func(logger *zap.Logger) echo.MiddlewareFunc {
		return middleware.BodyDumpWithConfig(
			middleware.BodyDumpConfig{
				Skipper: options.Skippers.Handler,
				Handler: func(ctx echo.Context, reqBody, resBody []byte) {
					logger.Info("dump body",
						zap.String("path", ctx.Path()),

						// request
						zap.Dict("request",
							zap.String("id", ctx.Response().Header().Get(echo.HeaderXRequestID)),
							zap.Any("headers", ctx.Request().Header),
							zap.ByteString("body", reqBody),
						),

						// response
						zap.Dict("response",
							zap.Any("headers", ctx.Response().Header()),
							zap.ByteString("body", resBody),
						),
					)
				},
			},
		)
	}
}
