package graphql

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/x-challenges/raven/common/errors"
	"go.uber.org/zap"
)

type recoveryFn = graphql.RecoverFunc

func newRecoveryFn(logger *zap.Logger) recoveryFn {
	return func(_ context.Context, err any) error {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr)
		debug.PrintStack()

		// TODO:

		logger.Error("graphql recovery",
			zap.Any("error", err),
		)

		e := gqlerror.Errorf("Internal server error")

		e.Extensions = map[string]interface{}{
			"code": errors.InternalError,
		}

		return e
	}
}
