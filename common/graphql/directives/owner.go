package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"

	"github.com/x-challenges/raven/common/errors"
)

// Ownerable
type Ownerable interface {
	GetOwner() string
}

// UserIDExtractor
type UserIDExtractor func(context.Context) string

// OwnerDirectiveFn
type OwnerDirectiveFn func(ctx context.Context, v interface{}, next graphql.Resolver) (interface{}, error)

// Owner
func Owner(userIDExtractor UserIDExtractor) OwnerDirectiveFn {
	return func(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
		var (
			res any
			err error
		)

		if res, err = next(ctx); err != nil {
			return nil, err
		}

		if obj, ok := res.(Ownerable); ok {
			if obj.GetOwner() == userIDExtractor(ctx) {
				return next(ctx)
			}
			return nil, errors.WithMessage(errors.ErrForbidden, "forbidden")
		}

		return nil, errors.WithMessage(errors.ErrInternal, "object dont implement owner interface")
	}
}
