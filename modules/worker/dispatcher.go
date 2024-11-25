package worker

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// Dispatcher interface.
type Dispatcher interface {
	// Run all workers
	Run(ctx context.Context) error

	// Shutdown all workers
	Shutdown(ctx context.Context) error

	// Register worker in dispatcher
	Register(worker Worker)
}

// Dispatcher interface implementations.
type dispatcher struct {
	logger  *zap.Logger
	workers []Worker
}

// NewDispatcher construct.
func NewDispatcher(logger *zap.Logger) Dispatcher {
	return &dispatcher{
		logger:  logger,
		workers: []Worker{},
	}
}

// Run implements Dispatcher interface.
func (d *dispatcher) Run(ctx context.Context) error {
	var (
		wg   sync.WaitGroup
		errs []error
	)

	for i := range d.workers {
		wg.Add(1)

		var w = d.workers[i]

		go func() {
			defer wg.Done()

			if err := w.Run(ctx); err != nil {
				errs = append(errs, fmt.Errorf("worker run failed: %w", err))
			}
		}()
	}
	wg.Wait()

	return errors.Join(errs...)
}

// Shutdown implements Dispatcher interface.
func (d *dispatcher) Shutdown(ctx context.Context) error {
	var (
		wg   sync.WaitGroup
		errs []error
	)

	for i := range d.workers {
		wg.Add(1)

		var w = d.workers[i]

		go func() {
			defer wg.Done()

			if err := w.Shutdown(ctx); err != nil {
				errs = append(errs, err)
			}
		}()
	}
	wg.Wait()

	return errors.Join(errs...)
}

// Register implements Dispatcher interface.
func (d *dispatcher) Register(worker Worker) {
	d.workers = append(d.workers, worker)
}
