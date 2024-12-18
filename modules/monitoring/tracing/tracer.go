package tracing

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
)

// InvokeTraceProvider
func InvokeTraceProvider(config *Config) error {
	if !config.Monitoring.Tracing.Enabled {
		return nil
	}

	var (
		spanExporter   trace.SpanExporter
		spanProcessor  trace.SpanProcessor
		tracerProvider *trace.TracerProvider
		err            error
	)

	if spanExporter, err = otlptracegrpc.New(context.Background()); err != nil {
		return err
	}

	// init span processor
	spanProcessor = trace.NewBatchSpanProcessor(spanExporter)

	// init trace provider
	tracerProvider = trace.NewTracerProvider(
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.Service.Name),
		)),
		trace.WithSpanProcessor(spanProcessor),
		trace.WithSampler(trace.AlwaysSample()),
	)

	// register trace provider
	otel.SetTracerProvider(tracerProvider)

	return nil
}
