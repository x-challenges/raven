package flood

import (
	"context"

	asynqpkg "github.com/hibiken/asynq"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/asynq"
	"github.com/x-challenges/raven/modules/config"
)

var ModuleName = "flood"

var Module = func(opts ...ModuleOption) fx.Option {
	var options = &ModuleOptions{}

	// load options
	for _, opt := range opts {
		opt(options)
	}

	// build and return fx module
	return fx.Module(
		ModuleName,

		config.Inject(new(Config)),

		// third party dependency
		asynq.Module(options.Server, options.Scheduler),

		// public usage
		fx.Provide(
			fx.Annotate(newClient,
				fx.As(new(Client)),
				fx.OnStop(func(ctx context.Context, client Client) error { return client.Close(ctx) }),
			),
		),

		fx.Decorate(
			func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
		),
	)
}

type ModuleOptions struct {
	Scheduler bool // enable scheduler instance in fx
	Server    bool // enable server instance in fx
}

type ModuleOption func(*ModuleOptions)

func WithScheduler(v bool) ModuleOption {
	return func(opts *ModuleOptions) {
		opts.Scheduler = v
	}
}

func WithServer(v bool) ModuleOption {
	return func(opts *ModuleOptions) {
		opts.Server = v
	}
}

// build and register worker to asynq
func Inject[W Worker[T], T JobArgs](builder interface{}) fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(builder, fx.As(new(Worker[T]))),
		),

		fx.Invoke(
			func(mux *asynqpkg.ServeMux, scheduler *asynqpkg.Scheduler, w Worker[T]) error {
				var (
					task    = newAsynqTask(w, []byte("{}"))
					options = newAsynqOptionFromWorker(w)
					handler = newAsynqHandlerFn(w)
					err     error
				)

				// register worker as a task handler
				mux.HandleFunc(task.Type(), handler)

				// register worker as a scheduler task
				if w, ok := w.(WorkerCron[JobArgs]); ok {
					if _, err = scheduler.Register(w.Cronspec(), task, options...); err != nil {
						return err
					}
				}

				return nil
			},
		),
	)
}
