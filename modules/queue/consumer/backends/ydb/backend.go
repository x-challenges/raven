package ydb

import (
	"context"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/topic/topicoptions"
	"github.com/ydb-platform/ydb-go-sdk/v3/topic/topicreader"
	"go.uber.org/zap"

	"github.com/x-challenges/raven/modules/queue/consumer/backends/backend"
)

const Type backend.Type = "YDB"

type Backend struct {
	logger *zap.Logger
	client *ydb.Driver
	reader *topicreader.Reader

	stop chan struct{}
	done chan struct{}
}

var _ backend.Backend = (*Backend)(nil)

func NewBackend(logger *zap.Logger, client *ydb.Driver, config *Config) (*Backend, error) {
	var (
		reader *topicreader.Reader
		err    error
	)

	var (
		readerOptions = []topicoptions.ReaderOption{
			topicoptions.WithReaderCommitMode(topicoptions.CommitModeSync),
		}

		selector = topicoptions.ReadTopic(config.Topic)
	)

	if reader, err = client.Topic().StartReader(config.Consumer, selector, readerOptions...); err != nil {
		return nil, err
	}

	return &Backend{
		logger: logger,
		client: client,
		reader: reader,
		stop:   make(chan struct{}),
		done:   make(chan struct{}),
	}, nil
}

// Run implements Backend interface
func (b *Backend) Run(ctx context.Context, callback backend.Callback) error {
	var err error

	// wait reader initialization or return error
	if err = b.reader.WaitInit(ctx); err != nil {
		return err
	}

	// async topic reader fn
	go func() {
		var (
			msg *topicreader.Message
			err error
		)

	loop:
		for {
			select {
			case <-b.stop:
				break loop
			default:
				// read message
				if msg, err = b.reader.ReadMessage(context.TODO()); err != nil {
					b.logger.Error("failed to read message", zap.Error(err))

					time.Sleep(5 * time.Second) // TODO: need some retries and fail app ...

					continue
				}

				// process message
				if err = callback(ctx, msg); err != nil {
					b.logger.Error("failed to process message, skip commit",
						zap.Any("msg", msg),
						zap.Error(err),
					)

					// skip commit message
					continue
				}

				// commit message
				if err = b.reader.Commit(ctx, msg); err != nil {
					b.logger.Error("failed to commit message, wait retry",
						zap.Any("msg", msg),
						zap.Error(err),
					)
				}
			}
		}

		// send done signal
		close(b.done)
	}()

	return nil
}

// Close implements Backend interface
func (b *Backend) Close(ctx context.Context) error {
	// send stop signal
	close(b.stop)

	// wait done signal
	<-b.done

	// close reader
	return b.reader.Close(ctx)
}
