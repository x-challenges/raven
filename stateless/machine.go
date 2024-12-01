package stateless

import (
	"context"

	"github.com/qmuntal/stateless"
)

// Args
type Args interface{}

// State
type State string

// Trigger
type Trigger string

// Machine
type Machine[T Args] struct {
	stateless *stateless.StateMachine
}

// New
func New[T Args](initial State) *Machine[T] {
	return &Machine[T]{
		stateless: stateless.NewStateMachine(initial),
	}
}

// Configure
func (sm *Machine[T]) Configure(state State) *StateConfiguration[T] {
	return &StateConfiguration[T]{
		stateless: sm.stateless.Configure(state),
	}
}

// Fire
func (sm *Machine[T]) Fire(ctx context.Context, trigger Trigger, args T) error {
	return sm.stateless.FireCtx(ctx, trigger, args)
}
