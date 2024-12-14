package metrics

import (
	"context"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
)

// ModuleName
var ModuleName = "metrics"

// Module
var Module = fx.Module(
	ModuleName,

	config.Inject(new(Config)),

	// public usage
	fx.Provide(
		fx.Annotate(
			NewExporter,

			fx.OnStop(func(ctx context.Context, provider *metric.MeterProvider) error {
				return provider.Shutdown(ctx)
			}),
		),
	),

	// register http controllers
	fx.Invoke(
		NewMetricsController,
	),

	// force
	fx.Invoke(func(*prometheus.Exporter) {}),

	fx.Decorate(
		func(logger *zap.Logger) *zap.Logger { return logger.Named(ModuleName) },
	),
)
