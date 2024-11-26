package trigger

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/yandex/trigger/messages"
)

// TimerMessages
type TimerMessages = messages.Messages[messages.Timer]

// TimerCallback
type TimerCallback func(context.Context, TimerMessages) error

// Timer
type Timer struct {
	logger       *zap.Logger
	callback     TimerCallback
	errorHandler ErrorHandler
}

var _ Trigger = (*Timer)(nil)

// ServeHTTP implements Trigger interface
func (t *Timer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	triggerHandler(&trigger[TimerMessages]{
		logger:       t.logger,
		builder:      func() TimerMessages { return TimerMessages{} },
		callback:     t.callback,
		errorHandler: t.errorHandler,
	}, w, r)
}
