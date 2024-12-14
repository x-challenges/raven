package metrics

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	sdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// NewExporter
func NewExporter() (*sdk.MeterProvider, *prometheus.Exporter, error) {
	var (
		exporter *prometheus.Exporter
		provider *sdk.MeterProvider
		err      error
	)

	// setup resources
	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		// semconv.ServiceNameKey.String(build.Service),
		// semconv.ServiceVersionKey.String(build.Version),
	)

	// setup prometheus exporter
	if exporter, err = prometheus.New(); err != nil {
		return nil, nil, err
	}

	// setup provider
	provider = sdk.NewMeterProvider(
		sdk.WithResource(resources),
		sdk.WithReader(exporter),
	)

	// setup otel with selected provider
	otel.SetMeterProvider(provider)

	return provider, exporter, nil
}
