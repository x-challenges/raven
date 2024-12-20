package tracing

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// FromSpan
func FromSpan(span trace.Span) zap.Field {
	var spanContext = span.SpanContext()

	return zap.Dict("tracing",
		zap.String("trace_id", spanContext.TraceID().String()),
		zap.String("span_id", spanContext.SpanID().String()),
	)
}

// FromContext
func FromContext(ctx context.Context) zap.Field {
	return FromSpan(trace.SpanFromContext(ctx))
}
