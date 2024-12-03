package asynq

import (
	"github.com/x-challenges/raven/modules/asynq/middlewares"
	"github.com/x-challenges/raven/modules/config"
	"github.com/x-challenges/raven/modules/worker"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const ModuleName = "asynq"

func Module(server, scheduler bool) fx.Option {
	var (
		fxOptions = []fx.Option{}
	)

	if server {
		fxOptions = append(fxOptions,
			fx.Provide(
				NewServer,
				NewServeMux,
			),

			fx.Provide(
				fx.Annotate(
					middlewares.NewLoggingMiddleware,
					fx.ResultTags(`group:"asynq:middlewares"`),
				),
			),

			worker.Inject[ServerWorker](NewServerWorker),
		)
	}

	if scheduler {
		fxOptions = append(fxOptions,
			fx.Provide(
				NewScheduler,
			),

			worker.Inject[SchedulerWorker](NewSchedulerWorker),
		)
	}

	return fx.Module(
		ModuleName,

		config.Inject(new(Config)),

		// public usage
		fx.Provide(
			NewClient,
			NewErrorHandlers,
		),

		fx.Options(fxOptions...),

		fx.Decorate(
			func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
		),
	)
}
