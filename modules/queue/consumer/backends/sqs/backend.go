package sqs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/queue/consumer/backends/backend"
)

const Type backend.Type = "SQS"

type Backend struct {
	logger *zap.Logger
	client *sqs.Client
	config *Config

	stop chan struct{}
	done chan struct{}
}

var _ backend.Backend = (*Backend)(nil)

func NewBackend(logger *zap.Logger, client *sqs.Client, config *Config) (*Backend, error) {
	return &Backend{
		logger: logger,
		client: client,
		config: config,

		stop: make(chan struct{}),
		done: make(chan struct{}),
	}, nil
}

// Run implements backend.Backend interface
func (b *Backend) Run(ctx context.Context, callback backend.Callback) error {
	go func() {
		for {
			received, err := b.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
				QueueUrl:            &b.config.URL,
				MaxNumberOfMessages: 1,  // TODO: move to config ...
				WaitTimeSeconds:     60, // TODO: move to config ...
			})

			if err == context.DeadlineExceeded {
				time.Sleep(5 * time.Second) // TODO: move to config ...
				continue
			}

			if err != nil {
				b.logger.Error("cant receive messages from sqs, retry", zap.Error(err))
				time.Sleep(5 * time.Second) // TODO: move to config ...
				continue
			}

			for _, msg := range received.Messages {
				if err := callback(ctx, msg); err != nil {
					b.logger.Error("process sqs message failed",
						zap.Any("msg", msg),
						zap.Error(err),
					)
				}
			}
		}
	}()
	return nil
}

// Close implements backend.Backend interface
func (b *Backend) Close(context.Context) error {
	return nil
}
