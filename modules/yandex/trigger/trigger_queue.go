package trigger

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/yandex/trigger/messages"
)

// QueueMessages
type QueueMessages = messages.Messages[messages.Queue]

// QueueCallback
type QueueCallback func(context.Context, QueueMessages) error

// Queue
type Queue struct {
	logger       *zap.Logger
	callback     QueueCallback
	errorHandler ErrorHandler
}

var _ Trigger = (*Queue)(nil)

// ServeHTTP implements Trigger interface
func (q *Queue) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	triggerHandler(&trigger[QueueMessages]{
		logger:       q.logger,
		builder:      func() QueueMessages { return QueueMessages{} },
		callback:     q.callback,
		errorHandler: q.errorHandler,
	}, w, r)
}
