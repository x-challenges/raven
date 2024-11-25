package directives

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/landrade/gqlgen-cache-control-plugin/cache"
	"github.com/mitchellh/hashstructure/v2"
	"go.uber.org/zap"
)

type CacheScope string

const (
	CachePublic  CacheScope = "PUBLIC"
	CachePrivate CacheScope = "PRIVATE"
)

type CacheDirectiveFn func(
	context.Context, interface{}, graphql.Resolver, int, string, CacheScope, bool, bool) (interface{}, error)

// Cacher
type cacher interface {
	// Get
	Get(ctx context.Context, key string, obj any) error

	// SetTTl
	SetWithTTL(ctx context.Context, key string, obj any, ttl time.Duration) error
}

type CacheSkipper func() bool

type CacheUserIDExtractor func(context.Context) string

type CacheTypeBundle[T any] struct {
	Factory func() T
	Cast    func(T any) any
}

type CacheTypeBuilder map[string]CacheTypeBundle[any]

func buildCacheControlKey(
	ctx context.Context,
	fc *graphql.FieldContext,
	scope CacheScope,
	userIDExtractor CacheUserIDExtractor,
) (string, error) {
	var (
		hash uint64
		err  error
	)

	if len(fc.Args) > 0 {
		if hash, err = hashstructure.Hash(fc.Args, hashstructure.FormatV2, nil); err != nil {
			return "", err
		}
	}

	var scopeSuffix = ":scope"

	switch scope {
	case CachePublic:
		scopeSuffix += ":public"
	case CachePrivate:
		scopeSuffix = ":private:" + userIDExtractor(ctx)
	}

	return strings.ToLower(
		fc.Object + ":" +
			fc.Field.Name + ":" +
			strconv.FormatUint(hash, 10) +
			scopeSuffix,
	), nil
}

func NewCache(
	logger *zap.Logger,
	c cacher,
	skipper CacheSkipper,
	userIDExtractor CacheUserIDExtractor,
	typeBuilder CacheTypeBuilder,
) CacheDirectiveFn {
	return func(
		ctx context.Context,
		_ interface{},
		next graphql.Resolver,
		maxAge int,
		model string,
		scope CacheScope,
		skip bool,
		control bool,
	) (interface{}, error) {
		var (
			key string
			ttl = time.Duration(maxAge) * time.Second
			err error
		)

		if skip || skipper() {
			return next(ctx)
		}

		// set cache control header
		if control {
			cache.SetHint(ctx, cache.Scope(scope), ttl)
		}

		// build cache key
		key, err = buildCacheControlKey(ctx,
			graphql.GetFieldContext(ctx),
			scope,
			userIDExtractor,
		)

		if err != nil {
			logger.Error("build cache key failed", zap.Error(err))
			return next(ctx)
		}

		var result = typeBuilder[model].Factory()

		// try get from cache
		if err = c.Get(ctx, key, &result); err == nil && result != nil {
			return typeBuilder[model].Cast(result), nil
		}

		// resolve data
		if result, err = next(ctx); err != nil {
			return nil, err
		}

		// try set to cache
		if err = c.SetWithTTL(ctx, key, &result, ttl); err != nil {
			logger.Error("cache set failed", zap.Error(err))
		}

		return result, nil
	}
}
