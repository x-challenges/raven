package monitoring

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/monitoring/sentry"
)

const ModuleName = "monitoring"

var Module = fx.Module(
	ModuleName,

	sentry.Module,

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
