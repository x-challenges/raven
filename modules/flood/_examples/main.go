package main

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/config"
	"github.com/x-challenges/raven/modules/flood"
	"github.com/x-challenges/raven/modules/logger"
	"github.com/x-challenges/raven/modules/redis"
	"github.com/x-challenges/raven/modules/worker"
)

func main() {
	app := fx.New(
		config.Module(),
		logger.Module,
		redis.Module,
		worker.Module,

		flood.Module(
			flood.WithScheduler(true), // enable scheduler
			flood.WithServer(true),    // enable server
		),

		flood.Inject[*TestWorker](NewTestWorker),

		// fx.Invoke(
		// 	func(client flood.Client) error {
		// 		_, err := client.Enqueue(context.Background(),
		// 			TestArgs{
		// 				Foo: "foo",
		// 				Bar: "foo",
		// 			},
		// 			flood.WithTaskID(uuid.NewString()),
		// 			flood.WithMaxRetry(10),
		// 		)
		// 		return err
		// 	},
		// ),
	)

	app.Run()
}

type TestArgs struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

var _ flood.JobArgs = (*TestArgs)(nil)

func (TestArgs) Kind() string { return "testovo" }

type TestWorker struct {
	logger *zap.Logger

	flood.WorkerDefaults[TestArgs]     // worker defaults
	flood.WorkerCronDefaults[TestArgs] // worker is periodic job
}

var _ flood.Worker[TestArgs] = (*TestWorker)(nil)

var _ flood.WorkerCron[TestArgs] = (*TestWorker)(nil)

func NewTestWorker(logger *zap.Logger) *TestWorker {
	return &TestWorker{
		logger: logger,
	}
}

func (w *TestWorker) Queue(*flood.Job[TestArgs]) string { return "test" }
func (w *TestWorker) Group(*flood.Job[TestArgs]) string { return "test" }
func (w *TestWorker) MaxRetry(*flood.Job[TestArgs]) int { return 1 }

func (w *TestWorker) Cronspec() string { return "@every 5s" }

func (w *TestWorker) Work(_ context.Context, _ *flood.Job[TestArgs]) error {
	return nil
}
