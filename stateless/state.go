package stateless

import (
	"github.com/qmuntal/stateless"
)

// StateConfiguration
type StateConfiguration[T Args] struct {
	stateless *stateless.StateConfiguration
}

// Permit
func (sc *StateConfiguration[T]) Permit(trigger Trigger, state State, guards ...GuardFunc[T]) *StateConfiguration[T] {
	sc.stateless.Permit(trigger, state, wrapGuardFuncs(guards...)...)
	return sc
}

// OnEntry
func (sc *StateConfiguration[T]) OnEntry(callback ActionFunc[T]) *StateConfiguration[T] {
	sc.stateless.OnEntry(wrapActionFunc(callback))
	return sc
}

// OnEntryFrom
func (sc *StateConfiguration[T]) OnEntryFrom(trigger Trigger, callback ActionFunc[T]) *StateConfiguration[T] {
	sc.stateless.OnEntryFrom(trigger, wrapActionFunc(callback))
	return sc
}

// OnExit
func (sc *StateConfiguration[T]) OnExit(callback ActionFunc[T]) *StateConfiguration[T] {
	sc.stateless.OnExit(wrapActionFunc(callback))
	return sc
}

// OnExitWith
func (sc *StateConfiguration[T]) OnExitWith(trigger Trigger, callback ActionFunc[T]) *StateConfiguration[T] {
	sc.stateless.OnExitWith(trigger, wrapActionFunc(callback))
	return sc
}
