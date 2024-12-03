package broadcaster

// Listener
type Listener[T any] interface {
	// Updates
	Updates() <-chan T

	// Close
	Close()
}

// Listener
type listener[T any] struct {
	msgCh       chan T
	broadcaster *broadcaster[T]
}

var _ Listener[any] = (*listener[any])(nil)

// NewListener
func NewListener[T any](broadcaster *broadcaster[T]) Listener[T] {
	return &listener[T]{
		msgCh:       make(chan T, 10), // buffered
		broadcaster: broadcaster,
	}
}

// Updates
func (l *listener[T]) Updates() <-chan T {
	return l.msgCh
}

// Close
func (l *listener[T]) Close() {
	l.broadcaster.unsubscribe(l.msgCh)
}
