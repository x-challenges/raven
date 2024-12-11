package broadcaster

import "context"

// Broadcaster
type Broadcaster[T any] interface {
	// Start
	Start(context.Context)

	// Stop
	Stop(context.Context)

	// Subscribe
	Subscribe() Listener[T]

	// Publish
	Publish(msg T)
}

// Broadcaster interface implementation
type broadcaster[T any] struct {
	stopCh    chan struct{}
	publishCh chan T
	subCh     chan chan T
	unsubCh   chan chan T
}

var _ Broadcaster[any] = (*broadcaster[any])(nil)

// New
func New[T any]() *broadcaster[T] {
	return &broadcaster[T]{
		stopCh:    make(chan struct{}),
		publishCh: make(chan T, 1),
		subCh:     make(chan chan T, 10),
		unsubCh:   make(chan chan T, 10),
	}
}

// Start
func (b *broadcaster[T]) Start(context.Context) {
	go func() {
		subs := map[chan T]struct{}{}

		for {
			select {
			case <-b.stopCh:
				return

			// subscribe
			case msgCh := <-b.subCh:
				subs[msgCh] = struct{}{}

			// unsubscribe
			case msgCh := <-b.unsubCh:
				delete(subs, msgCh)

			// broadcast
			case msg := <-b.publishCh:
				for msgCh := range subs {
					// msgCh is buffered, use non-blocking send to protect the broker:
					select {
					case msgCh <- msg:
					default:
					}
				}
			}
		}
	}()
}

// Stop
func (b *broadcaster[T]) Stop(context.Context) {
	close(b.stopCh)
}

// Subscribe
func (b *broadcaster[T]) Subscribe() Listener[T] {
	var ln = NewListener(b)

	// subscribe
	b.subCh <- ln.(*listener[T]).msgCh

	return ln
}

// Unsubscribe
func (b *broadcaster[T]) unsubscribe(msgCh chan T) {
	b.unsubCh <- msgCh
}

// Publish
func (b *broadcaster[T]) Publish(msg T) {
	b.publishCh <- msg
}
