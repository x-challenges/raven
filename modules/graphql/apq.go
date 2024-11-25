package graphql

import (
	"context"

	"github.com/x-challenges/raven/modules/cache"
)

type apqCacheAdapter struct {
	cache.Cache[string]
}

func newAPQCacheAdapter(factory *cache.Factory, config *Config) (*apqCacheAdapter, error) {
	cache, err := cache.New[string](factory, &config.Graphql.Caches.APQ)
	if err != nil {
		return nil, err
	}

	return &apqCacheAdapter{
		Cache: cache,
	}, nil
}

// Add
func (a *apqCacheAdapter) Add(ctx context.Context, key string, obj string) {
	_ = a.Cache.Set(ctx, key, obj)
}

// Get
func (a *apqCacheAdapter) Get(ctx context.Context, key string) (string, bool) {
	var (
		res string
		err error
	)

	if res, err = a.Cache.Get(ctx, key); err != nil {
		return "", false
	}

	return res, true
}
