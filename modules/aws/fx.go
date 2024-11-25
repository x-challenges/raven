package aws

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
)

const ModuleName = "aws"

var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	fx.Provide(
		NewAwsConfig,
		NewClient,
	),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
