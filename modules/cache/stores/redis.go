package stores

import (
	"context"
	"time"

	lib_store "github.com/eko/gocache/lib/v4/store"
	store "github.com/eko/gocache/store/redis/v4"
	goredis "github.com/redis/go-redis/v9"
)

const Redis kind = "redis"

type redisOptions struct {
	Prefix string        `mapstructure:"prefix" default:"cache"`
	TTL    time.Duration `mapstructure:"ttl" default:"60s"`
}

type redis struct {
	*store.RedisStore
	options *redisOptions
}

var _ Store = (*redis)(nil)

func newRedisFactory(client goredis.UniversalClient) factory {
	return func() (kind, builder) {
		return Redis, func(options *Options) (Store, error) {
			return newRedis(client, options.Redis)
		}
	}
}

func newRedis(client goredis.UniversalClient, options *redisOptions) (*redis, error) {
	var (
		opts = []lib_store.Option{
			lib_store.WithExpiration(options.TTL),
			lib_store.WithTags([]string{options.Prefix}),
		}
	)

	return &redis{
		RedisStore: store.NewRedis(client, opts...),
		options:    options,
	}, nil
}

// Clear implements Store interface
func (r *redis) Clear(ctx context.Context) error {
	return r.Invalidate(ctx, lib_store.WithInvalidateTags([]string{r.options.Prefix}))
}
