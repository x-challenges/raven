package cache

import (
	"context"
	"errors"
	"time"

	gocache "github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"

	"github.com/x-challenges/raven/modules/cache/stores"
)

type Options = stores.Options

type (
	Key    = string
	Object = any
)

type Cache[T Object] interface {
	// Get
	Get(context.Context, Key) (T, error)

	// GetWithTTL
	GetWithTTL(context.Context, Key) (T, time.Duration, error)

	// Set
	Set(context.Context, Key, T) error

	// SetWithTTL
	SetWithTTL(context.Context, Key, T, time.Duration) error

	// Delete
	Delete(context.Context, Key) error

	// Clear
	Clear(ctx context.Context) error

	// Try
	Try(context.Context, Key, func(context.Context, Key) (T, error)) (T, error)
}

type cache[T Object] struct {
	manager *gocache.Cache[T]
}

var _ Cache[Object] = (*cache[Object])(nil)

type Factory = stores.Factory

func New[T Object](factory *Factory, options *Options) (Cache[T], error) {
	var (
		store stores.Store
		err   error
	)

	// build store
	if store, err = factory.Store(options); err != nil {
		return nil, err
	}

	return &cache[T]{
		manager: gocache.New[T](store),
	}, nil
}

// Get implements Cache interface
func (c *cache[T]) Get(ctx context.Context, key Key) (T, error) {
	obj, err := c.manager.Get(ctx, key)
	return obj, err
}

// GetWithTTL implements Cache interface
func (c *cache[T]) GetWithTTL(ctx context.Context, key Key) (T, time.Duration, error) {
	return c.manager.GetWithTTL(ctx, key)
}

// Set implements Cache interface
func (c *cache[T]) Set(ctx context.Context, key Key, obj T) error {
	return c.manager.Set(ctx, key, obj)
}

// SetWithTTL implements Cache interface
func (c *cache[T]) SetWithTTL(ctx context.Context, key Key, obj T, ttl time.Duration) error {
	return c.manager.Set(ctx, key, obj, store.WithExpiration(ttl))
}

// Delete implements Cache interface
func (c *cache[T]) Delete(ctx context.Context, key Key) error {
	return c.manager.Delete(ctx, key)
}

// Clear implements Cache interface
func (c *cache[T]) Clear(ctx context.Context) error {
	return c.manager.Clear(ctx)
}

// Try implements Cache interface
func (c *cache[T]) Try(ctx context.Context, key Key, call func(context.Context, Key) (T, error)) (T, error) {
	var (
		res T
		err error
	)

	// try get cache
	if res, err = c.Get(ctx, key); err != nil {
		if !errors.Is(err, &store.NotFound{}) {
			return res, err
		}
	}

	// exec callback
	if res, err = call(ctx, key); err != nil {
		return res, err
	}

	// set result to cache
	return res, c.Set(ctx, key, res)
}
