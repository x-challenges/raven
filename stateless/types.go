package stateless

import (
	"context"

	"github.com/qmuntal/stateless"
)

// GuardFunc
type GuardFunc[T Args] func(ctx context.Context, args T) bool

// ActionFunc
type ActionFunc[T Args] func(ctx context.Context, args T) error

// WrapGuardFunc
func wrapGuardFunc[T Args](guard GuardFunc[T]) stateless.GuardFunc {
	return func(ctx context.Context, args ...any) bool {
		return guard(ctx, args[0].(T))
	}
}

// WrapGuardFuncs
func wrapGuardFuncs[T Args](guards ...GuardFunc[T]) []stateless.GuardFunc {
	var result = make([]stateless.GuardFunc, len(guards))

	for _, guard := range guards {
		result = append(result, wrapGuardFunc(guard))
	}

	return result
}

// WrapActionFunc
func wrapActionFunc[T Args](action ActionFunc[T]) stateless.ActionFunc {
	return func(ctx context.Context, args ...any) error {
		return action(ctx, args[0].(T))
	}
}
