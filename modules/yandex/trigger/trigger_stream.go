package trigger

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

// StreamMessages
type StreamMessages struct {
	Messages []any `json:"messages"`
}

// StreamCallback
type StreamCallback func(context.Context, StreamMessages) error

// Stream
type Stream struct {
	logger       *zap.Logger
	callback     StreamCallback
	errorHandler ErrorHandler
}

var _ Trigger = (*Stream)(nil)

// ServeHTTP implements Trigger interface
func (s *Stream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	triggerHandler(&trigger[StreamMessages]{
		logger:       s.logger,
		builder:      func() StreamMessages { return StreamMessages{} },
		callback:     s.callback,
		errorHandler: s.errorHandler,
	}, w, r)
}
