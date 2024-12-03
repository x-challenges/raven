package flood

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// WorkerOptions
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

	// ErrorHandler
	ErrorHandler(context.Context, *Job[T], error)
}

// Worker
type Worker[T JobArgs] interface {
	WorkerOptions[T]

	// Work, need client impleemntation
	Work(ctx context.Context, job *Job[T]) error
}

// WorkerDefaults implements Worker interface
type WorkerDefaults[T JobArgs] struct{}

var _ WorkerOptions[JobArgs] = (*WorkerDefaults[JobArgs])(nil)

func (WorkerDefaults[T]) TaskID(*Job[T]) string                        { return uuid.NewString() }
func (WorkerDefaults[T]) Queue(*Job[T]) string                         { return "default" }
func (WorkerDefaults[T]) Group(*Job[T]) string                         { return "" }
func (WorkerDefaults[T]) Unique(*Job[T]) time.Duration                 { return 0 * time.Second }
func (WorkerDefaults[T]) Timeout(*Job[T]) time.Duration                { return 0 * time.Second }
func (WorkerDefaults[T]) MaxRetry(*Job[T]) int                         { return 10 }
func (WorkerDefaults[T]) ProcessIn(*Job[T]) time.Duration              { return 0 * time.Second }
func (WorkerDefaults[T]) ProcessAt(*Job[T]) time.Time                  { return time.Time{} }
func (WorkerDefaults[T]) Deadline(*Job[T]) time.Time                   { return time.Time{} }
func (WorkerDefaults[T]) Retention(*Job[T]) time.Duration              { return 0 * time.Second }
func (WorkerDefaults[T]) ErrorHandler(context.Context, *Job[T], error) {}

// WorkerCron
type WorkerCron[T JobArgs] interface {
	Cronspec() string
}

// WorkerCronDefaults
type WorkerCronDefaults[T JobArgs] struct{}

// Cronspec implements WorkerCron interface
func (WorkerCronDefaults[T]) Cronspec() string { return "" }

// WorkerWrapper
type WorkerWrapper[T JobArgs] struct {
	worker Worker[T]
}
