package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"

	"github.com/x-challenges/raven/common/validate"
)

func Constraint(ctx context.Context, _ interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	var (
		val interface{}
		err error
	)

	if val, err = next(ctx); err != nil {
		return nil, err
	}

	if err = validate.Var(val, constraint); err != nil {
		return nil, err // TODO:
	}
	return val, nil
}
