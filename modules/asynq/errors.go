package asynq

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

// ErrorHandler
type ErrorHandler = asynq.ErrorHandler

// ErrorHandlerFunc
type ErrorHandlerFunc = asynq.ErrorHandlerFunc

// ErrorHandlers
type ErrorHandlers struct {
	logger   *zap.Logger
	Handlers map[string]ErrorHandlerFunc
}

// NewErrorHandlers
func NewErrorHandlers(logger *zap.Logger) *ErrorHandlers {
	return &ErrorHandlers{
		logger:   logger,
		Handlers: make(map[string]asynq.ErrorHandlerFunc),
	}
}

// Register
func (e *ErrorHandlers) Register(task *asynq.Task, fn ErrorHandlerFunc) error {
	var (
		exist bool
	)

	// raise error if handler already exist
	if _, exist = e.Handlers[task.Type()]; exist {
		return fmt.Errorf("asynq task already registered,%s", task.Type())
	}

	// register error handler
	e.Handlers[task.Type()] = fn

	return nil
}

// HandleError
func (e *ErrorHandlers) HandleError(ctx context.Context, task *asynq.Task, err error) {
	var (
		handler ErrorHandlerFunc
		exist   bool
	)

	// skip if error handler not exist
	if handler, exist = e.Handlers[task.Type()]; !exist {
		return
	}

	// handle error
	handler(ctx, task, err)
}
