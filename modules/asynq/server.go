package asynq

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/worker"
)

func NewServer(logger *zap.Logger, client redis.UniversalClient, config *Config) (*asynq.Server, error) {
	var (
		opts = asynq.Config{
			Logger:          logger.Sugar(),
			Concurrency:     config.Asynq.Server.Concurrency,
			Queues:          config.Asynq.Server.Queues,
			StrictPriority:  config.Asynq.Server.StrictPriority,
			ShutdownTimeout: config.Asynq.Server.Shutdown.Timeout,
		}
	)

	return asynq.NewServerFromRedisClient(client, opts), nil
}

type ServeMuxParams struct {
	fx.In
	Middlewares []asynq.MiddlewareFunc `group:"asynq:middlewares"`
}

// NewServeMux construct.
func NewServeMux(p ServeMuxParams) (*asynq.ServeMux, error) {
	var (
		mux = asynq.NewServeMux()
	)

	// register middlewares
	for _, middleware := range p.Middlewares {
		mux.Use(middleware)
	}

	return mux, nil
}

type ServerWorker interface {
	worker.Worker
}

type serverWorker struct {
	server *asynq.Server
	mux    *asynq.ServeMux
}

func NewServerWorker(server *asynq.Server, mux *asynq.ServeMux) ServerWorker {
	return &serverWorker{
		server: server,
		mux:    mux,
	}
}

// Run implements worker.Worker interface.
func (s *serverWorker) Run(_ context.Context) error {
	return s.server.Start(s.mux)
}

// Shutdown implements worker.Worker interface.
func (s *serverWorker) Shutdown(_ context.Context) error {
	s.server.Shutdown()
	return nil
}
