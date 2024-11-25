package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ModuleName.
const ModuleName = "config"

// Module provided to fx.
var Module = func(opts ...Option) fx.Option {
	options := &options{}

	// load functional options
	for _, opt := range opts {
		opt(options)
	}

	return fx.Module(
		ModuleName,

		fx.Provide(
			viper.New,
			NewLoader(options.optionalPrefix),
		),
	)
}

// Inject config to fx.
func Inject[C Config](c *C) fx.Option {
	return fx.Options(
		fx.Supply(c),

		fx.Decorate(
			func(logger *zap.Logger, loader Loader, config *C) *C {
				if err := loader.Load(config); err != nil {
					logger.Fatal("config load failed", zap.Error(err))
				}
				return config
			},
		),
	)
}
