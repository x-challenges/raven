package main

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/aws"
	"github.com/x-challenges/raven/modules/config"
	"github.com/x-challenges/raven/modules/logger"

	"github.com/x-challenges/raven/modules/queue/consumer"
)

const TestConsumerName = "test"

type TestConsumer = consumer.Consumer

type testConsumer struct {
	logger *zap.Logger
}

var _ TestConsumer = (*testConsumer)(nil)

func NewTestConsumer(logger *zap.Logger) (TestConsumer, error) {
	return &testConsumer{
		logger: logger,
	}, nil
}

// Config implements consumer.Consumer interface
func (testConsumer) Name() string { return TestConsumerName }

// Process implements consumer.Consumer interface
func (tc *testConsumer) Process(_ context.Context, messages ...consumer.Message) error {
	tc.logger.Debug("received message",
		zap.Any("messages", messages),
	)
	return nil
}

func main() {
	config.Files = append(config.Files, "config.yaml")

	app := fx.New(
		config.Module(),
		logger.Module,
		aws.Module,
		consumer.Module,

		consumer.Inject[TestConsumer](NewTestConsumer),
	)

	app.Run()
}
