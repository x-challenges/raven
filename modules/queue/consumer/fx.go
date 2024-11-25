package consumer

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
	"github.com/x-challenges/raven/modules/queue/consumer/backends"
)

const ModuleName = "consumer"

var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	backends.Module,

	fx.Provide(
		fx.Annotate(
			NewDispatcher,
			fx.OnStart(func(ctx context.Context, d Dispatcher) error { return d.Run(ctx) }),
			fx.OnStop(func(ctx context.Context, d Dispatcher) error { return d.Shutdown(ctx) }),
		),
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),

	// fx.Invoke(
	// 	func(config *Config) error {
	// 		spew.Dump(config)
	// 		return nil
	// 	},
	// ),
)

func Inject[T Consumer](constructor interface{}) fx.Option {
	return fx.Options(
		fx.Provide(fx.Private,
			fx.Annotate(constructor, fx.As(new(T))),
		),

		// register consumer to dispatcher
		fx.Invoke(
			func(d Dispatcher, c T) error { return d.Register(c) },
		),
	)
}
