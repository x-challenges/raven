package main

import (
	"context"

	"go.uber.org/fx"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/x-challenges/raven/modules/cache"
	"github.com/x-challenges/raven/modules/config"
	"github.com/x-challenges/raven/modules/logger"
	"github.com/x-challenges/raven/modules/redis"
)

type User struct{ ID string }

type Cache cache.Cache[*User]

func NewCache(factory *cache.Factory) (Cache, error) {
	return cache.New[*User](factory, nil)
}

func main() {
	var app = fx.New(
		config.Module(),
		logger.Module,
		redis.Module,
		cache.Module,

		fx.Provide(
			NewCache,
		),

		fx.Invoke(
			func(cache Cache) {
				var (
					ctx  = context.TODO()
					user = &User{
						ID: uuid.NewString(),
					}
				)

				spew.Dump(cache.Get(ctx, user.ID))
				spew.Dump(cache.Set(ctx, user.ID, user))
				spew.Dump(cache.Try(ctx, user.ID, func(context.Context, string) (*User, error) { return user, nil }))

			},
		),
	)

	app.Run()
}
