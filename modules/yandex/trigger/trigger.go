package trigger

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// Trigger
type Trigger interface {
	// ServeHTTP
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// ErrorHandler
type ErrorHandler func(context.Context, error) error

// DefaultErrorHandler
var DefaultErrorHandler = func(_ context.Context, err error) error { return err }

// Trigger
type trigger[T any] struct {
	logger       *zap.Logger
	builder      func() T
	callback     func(context.Context, T) error
	errorHandler ErrorHandler
}

// TriggerHandler
func triggerHandler[T any](t *trigger[T], w http.ResponseWriter, r *http.Request) {
	var (
		logger = t.logger
		msgs   = t.builder()
		err    error
	)

	defer func() {
		if err = t.errorHandler(r.Context(), err); err == nil {
			return
		}

		logger.Error("trigger process failed", zap.Error(err))

		http.Error(w, "trigger process failed", http.StatusInternalServerError)
	}()

	// decode json
	if err = json.NewDecoder(r.Body).Decode(&msgs); err != nil {
		return
	}

	logger = logger.With(zap.Any("messages", msgs))

	if err = t.callback(r.Context(), msgs); err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
}
