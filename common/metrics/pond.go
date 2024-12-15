package metrics

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	// PoolMeter
	poolMeter = otel.Meter("pond/pool")
)

// BasePool
type poolBase interface {
	RunningWorkers() int64
	SubmittedTasks() uint64
	WaitingTasks() uint64
	FailedTasks() uint64
	SuccessfulTasks() uint64
	CompletedTasks() uint64
}

// InitPoolMetrics
func InitPoolMetrics(namespace string, pool poolBase) error {
	var (
		attributes = []attribute.KeyValue{
			attribute.String("namespace", namespace),
		}
		err error
	)

	// WorkersRunning
	if _, err = poolMeter.Int64ObservableGauge("pond.workers.running",
		metric.WithDescription("running workers"),
		metric.WithUnit("{worker}"),
		metric.WithInt64Callback(
			func(_ context.Context, io metric.Int64Observer) error {
				io.Observe(
					pool.RunningWorkers(),
					metric.WithAttributes(attributes...),
				)
				return nil
			},
		),
	); err != nil {
		return err
	}

	// TasksSubmittedTotal
	if _, err = poolMeter.Int64ObservableGauge("pond.tasks.submitted.total",
		metric.WithDescription("tasks submitted total"),
		metric.WithUnit("{task}"),
		metric.WithInt64Callback(
			func(_ context.Context, io metric.Int64Observer) error {
				io.Observe(
					int64(pool.SubmittedTasks()), // nolint:gosec
					metric.WithAttributes(attributes...),
				)
				return nil
			},
		),
	); err != nil {
		return err
	}

	// TasksWaitingTotal
	if _, err = poolMeter.Int64ObservableGauge("pond.tasks.waiting.total",
		metric.WithDescription("tasks waiting total"),
		metric.WithUnit("{task}"),
		metric.WithInt64Callback(
			func(_ context.Context, io metric.Int64Observer) error {
				io.Observe(
					int64(pool.WaitingTasks()), // nolint:gosec
					metric.WithAttributes(attributes...),
				)
				return nil
			},
		),
	); err != nil {
		return err
	}

	// TasksSuccessfulTotal
	if _, err = poolMeter.Int64ObservableGauge("pond.tasks.successful.total",
		metric.WithDescription("tasks successful total"),
		metric.WithUnit("{task}"),
		metric.WithInt64Callback(
			func(_ context.Context, io metric.Int64Observer) error {
				io.Observe(
					int64(pool.SuccessfulTasks()), // nolint:gosec
					metric.WithAttributes(attributes...),
				)
				return nil
			},
		),
	); err != nil {
		return err
	}

	// TasksFailedTotal
	if _, err = poolMeter.Int64ObservableGauge("pond.tasks.failed.total",
		metric.WithDescription("tasks failed total"),
		metric.WithUnit("{task}"),
		metric.WithInt64Callback(
			func(_ context.Context, io metric.Int64Observer) error {
				io.Observe(
					int64(pool.FailedTasks()), // nolint:gosec
					metric.WithAttributes(attributes...),
				)
				return nil
			},
		),
	); err != nil {
		return err
	}

	// TasksCompletedTotal
	if _, err = poolMeter.Int64ObservableGauge("pond.tasks.completed.total",
		metric.WithDescription("tasks completed total"),
		metric.WithUnit("{task}"),
		metric.WithInt64Callback(
			func(_ context.Context, io metric.Int64Observer) error {
				io.Observe(
					int64(pool.CompletedTasks()), // nolint:gosec
					metric.WithAttributes(attributes...),
				)
				return nil
			},
		),
	); err != nil {
		return err
	}

	return nil
}
