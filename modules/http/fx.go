package http

import (
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
	"github.com/x-challenges/raven/modules/http/controllers"
	"github.com/x-challenges/raven/modules/http/middlewares"
	"github.com/x-challenges/raven/modules/worker"
)

const ModuleName = "http"

// Module provided to fx.
var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	// server
	fx.Provide(
		NewServer,
		NewServerMux,
	),

	// middlewares
	fx.Provide(
		AsMiddleware(middleware.Recover),
		AsMiddleware(middleware.RequestID),
		AsMiddleware(middleware.CORS),

		// custom
		AsMiddleware(middlewares.LoggerMiddleware()),
		// AsMiddleware(middlewares.DumpMiddleware()),
		// AsMiddleware(middlewares.MetricsMiddleware()),

	),

	worker.Inject[ServerWorker](NewServerWorker),

	fx.Invoke(
		controllers.NewHealthCheckController,
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)

var AsMiddleware = func(middleware any) any {
	return fx.Annotate(
		middleware,
		fx.ResultTags(`group:"http:middlewares"`),
	)
}
