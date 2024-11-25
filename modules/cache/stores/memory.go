package stores

import (
	"time"

	gocache_store "github.com/eko/gocache/store/go_cache/v4"
	gocache "github.com/patrickmn/go-cache"
)

const Memory kind = "memory"

type memoryOptions struct {
	TTL     time.Duration `mapstructure:"ttl" default:"60s"`
	Cleanup time.Duration `mapstructure:"cleanup" default:"5m"`
}

type memory struct {
	*gocache_store.GoCacheStore
}

var _ Store = (*memory)(nil)

func newMemoryFactory() factory {
	return func() (kind, builder) {
		return Memory, func(options *Options) (Store, error) {
			return newMemory(options.Memory)
		}
	}
}

func newMemory(options *memoryOptions) (Store, error) {
	return &memory{
		GoCacheStore: gocache_store.NewGoCache(
			gocache.New(
				options.TTL,
				options.Cleanup,
			),
		),
	}, nil
}
