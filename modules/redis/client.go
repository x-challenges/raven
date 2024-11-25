package redis

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewClient(config *Config, logger *zap.Logger) (redis.UniversalClient, error) {
	var opts = &redis.UniversalOptions{
		Addrs:      config.Redis.Addresses,
		DB:         config.Redis.Database,
		ClientName: config.Redis.Client.Name,

		// policy
		DialTimeout:  config.Redis.Dial.Timeout,
		ReadTimeout:  config.Redis.Read.Timeout,
		WriteTimeout: config.Redis.Write.Timeout,

		// retry policy
		MaxRetries:      config.Redis.Retry.Count,
		MinRetryBackoff: config.Redis.Retry.MinBackoff,
		MaxRetryBackoff: config.Redis.Retry.MaxBackoff,

		// pool options
		PoolSize:    config.Redis.Pool.Size,
		PoolTimeout: config.Redis.Pool.Timeout,

		// conns
		MinIdleConns:    config.Redis.Pool.MinIdleConns,
		MaxIdleConns:    config.Redis.Pool.MaxIdleConns,
		MaxActiveConns:  config.Redis.Pool.MaxActiveConns,
		ConnMaxLifetime: config.Redis.Pool.ConnMaxLifetime,
	}

	// set logger
	redis.SetLogger(NewLoggerAdapter(logger))

	return redis.NewUniversalClient(opts), nil
}
