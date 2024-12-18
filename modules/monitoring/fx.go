package monitoring

import (
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/monitoring/metrics"
	"github.com/x-challenges/raven/modules/monitoring/sentry"
	"github.com/x-challenges/raven/modules/monitoring/tracing"
)

// ModuleName
const ModuleName = "monitoring"

// Module
var Module = fx.Module(
	ModuleName,

	sentry.Module,
	metrics.Module,
	tracing.Module,

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
