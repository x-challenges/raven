package asynq

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/worker"
)

// NewScheduler construct.
func NewScheduler(logger *zap.Logger, client redis.UniversalClient, config *Config) (*asynq.Scheduler, error) {
	var (
		loc *time.Location
		err error
	)

	if loc, err = time.LoadLocation(config.Asynq.Scheduler.Location); err != nil {
		return nil, err
	}

	return asynq.NewSchedulerFromRedisClient(client, &asynq.SchedulerOpts{
		Logger:   logger.Sugar(),
		Location: loc,
	}), nil
}

type SchedulerWorker interface {
	worker.Worker
}

type schedulerWorker struct {
	logger    *zap.Logger
	scheduler *asynq.Scheduler
}

type SchedulerWorkerParams struct {
	fx.In

	Logger    *zap.Logger
	Config    *Config
	Scheduler *asynq.Scheduler
}

// NewSchedulerWorker construct.
func NewSchedulerWorker(p SchedulerWorkerParams) (SchedulerWorker, error) {
	return &schedulerWorker{
		logger:    p.Logger,
		scheduler: p.Scheduler,
	}, nil
}

func (s *schedulerWorker) Run(_ context.Context) error {
	return s.scheduler.Start()
}

func (s *schedulerWorker) Shutdown(_ context.Context) error {
	s.scheduler.Shutdown()
	return nil
}
