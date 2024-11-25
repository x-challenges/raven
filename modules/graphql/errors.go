package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/common/errors"
)

type errorPresenterFn = graphql.ErrorPresenterFunc

func newErrorPresenterFn(logger *zap.Logger, config *Config) errorPresenterFn {
	return func(ctx context.Context, e error) *gqlerror.Error {
		var (
			err     = graphql.DefaultErrorPresenter(ctx, e)
			myError *errors.Error
		)

		if errors.As(err, &myError) {
			var (
				message    = errors.GetMessage(e)
				code       = errors.GetCode(e)
				level      = errors.GetLevel(e)
				fields     = errors.GetFields(e, config.Graphql.Errors.Debug)
				extensions = map[string]interface{}{}
			)

			extensions["code"] = code

			if !level.IsSystem() || config.Graphql.Errors.Debug {
				if len(fields) > 0 {
					extensions["details"] = fields.Values()
				}
			}

			if config.Graphql.Errors.Debug {
				extensions["level"] = level
			}

			err.Message = message
			err.Extensions = extensions

			logger.Log(level.ZapLogLevel(), "graphql error presenter",
				zap.Error(e),
			)
		}

		return err
	}
}
