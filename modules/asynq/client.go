package asynq

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

func NewClient(client redis.UniversalClient) *asynq.Client {
	return asynq.NewClientFromRedisClient(client)
}
