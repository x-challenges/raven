package limiter

import (
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

// New
func New(client redis.UniversalClient) *redis_rate.Limiter {
	return redis_rate.NewLimiter(client)
}
