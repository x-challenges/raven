package directives

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-redis/redis_rate/v10"

	"github.com/x-challenges/raven/common/errors"
)

// LimitScope
type LimitScope = string

const (
	LimitScopeUser LimitScope = "USER"
)

// LimitPeriod
type LimitPeriod = string

const (
	LimitPeriodSecond LimitPeriod = "SECOND"
	LimitPeriodMinute LimitPeriod = "MINUTE"
)

// Limiter
type Limiter func(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
	rate int,
	burst int,
	per LimitPeriod,
	scope LimitScope,
	skip bool,
) (interface{}, error)

type IPExtractor func(ctx context.Context) string

// NewLimiter
func NewLimiter(limiter *redis_rate.Limiter, ip IPExtractor, skipper func() bool) Limiter {
	return func(
		ctx context.Context, _ interface{}, next graphql.Resolver,
		rate, burst int, per LimitPeriod, scope LimitScope, skip bool,
	) (interface{}, error) {
		var (
			fc         = graphql.GetFieldContext(ctx)
			key string = fc.Object + ":" + fc.Field.Name
			res *redis_rate.Result
			err error
		)

		if skip || skipper() {
			return next(ctx)
		}

		// added user ip for private scope
		if scope == LimitScopeUser {
			key = key + ":ip:" + ip(ctx)
		}

		var lim = redis_rate.Limit{
			Rate:  rate,
			Burst: burst,
		}

		switch per {
		case LimitPeriodMinute:
			lim.Period = time.Minute
		default:
			lim.Period = time.Second
		}

		// check allow
		if res, err = limiter.Allow(ctx, key, lim); err != nil {
			return nil, err
		}

		// request not allowed
		if res.Allowed <= 0 {
			return nil, errors.WithFields(errors.ErrToManyRequest,
				errors.Int("remaining", res.Remaining),
				errors.Int("retry", int(res.RetryAfter.Seconds())),
			)
		}

		return next(ctx)
	}
}
