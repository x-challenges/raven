package flood

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type WorkerOptions[T JobArgs] interface {
	// TaskID
	TaskID(_ *Job[T]) string

	// Queue
	Queue(_ *Job[T]) string

	// Group
	Group(_ *Job[T]) string

	// Unique
	Unique(_ *Job[T]) time.Duration

	// Timeout
	Timeout(_ *Job[T]) time.Duration

	// MaxRetry
	MaxRetry(_ *Job[T]) int

	// ProcessIn
	ProcessIn(_ *Job[T]) time.Duration

	// ProcessAt
	ProcessAt(_ *Job[T]) time.Time

	// Deadline
	Deadline(_ *Job[T]) time.Time

	// Retention
	Retention(_ *Job[T]) time.Duration
}

type Worker[T JobArgs] interface {
	WorkerOptions[T]

	// Work, need client impleemntation
	Work(ctx context.Context, job *Job[T]) error
}

type WorkerDefaults[T JobArgs] struct{}

var _ WorkerOptions[JobArgs] = (*WorkerDefaults[JobArgs])(nil)

func (WorkerDefaults[T]) TaskID(_ *Job[T]) string           { return uuid.NewString() }
func (WorkerDefaults[T]) Queue(_ *Job[T]) string            { return "default" }
func (WorkerDefaults[T]) Group(_ *Job[T]) string            { return "" }
func (WorkerDefaults[T]) Unique(_ *Job[T]) time.Duration    { return 0 * time.Second }
func (WorkerDefaults[T]) Timeout(_ *Job[T]) time.Duration   { return 0 * time.Second }
func (WorkerDefaults[T]) MaxRetry(_ *Job[T]) int            { return 10 }
func (WorkerDefaults[T]) ProcessIn(_ *Job[T]) time.Duration { return 0 * time.Second }
func (WorkerDefaults[T]) ProcessAt(_ *Job[T]) time.Time     { return time.Time{} }
func (WorkerDefaults[T]) Deadline(_ *Job[T]) time.Time      { return time.Time{} }
func (WorkerDefaults[T]) Retention(_ *Job[T]) time.Duration { return 0 * time.Second }

type WorkerCron[T JobArgs] interface {
	Cronspec() string
}

type WorkerCronDefaults[T JobArgs] struct{}

func (WorkerCronDefaults[T]) Cronspec() string { return "" }

type WorkerWrapper[T JobArgs] struct {
	worker Worker[T]
}
