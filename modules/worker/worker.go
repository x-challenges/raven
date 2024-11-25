package worker

import "context"

// Worker interface.
type Worker interface {
	// Run worker
	Run(ctx context.Context) error

	// Shutdown worker
	Shutdown(ctx context.Context) error
}
