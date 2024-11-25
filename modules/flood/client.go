package flood

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

// Client
type Client interface {
	// Enqueue
	Enqueue(ctx context.Context, args JobArgs, opts ...Option) (*Result[JobArgs], error)

	// Close
	Close(ctx context.Context) error
}

// Client interface implementation
type client struct {
	logger *zap.Logger
	client *asynq.Client
}

// newClient
func newClient(logger *zap.Logger, asynq *asynq.Client) *client {
	return &client{
		logger: logger,
		client: asynq,
	}
}

// Enqueue implements Client interface
func (c *client) Enqueue(ctx context.Context, args JobArgs, opts ...Option) (*Result[JobArgs], error) {
	var (
		options = Options{}
		payload []byte
		err     error
	)

	for _, opt := range opts {
		opt(&options)
	}

	var wrappedArgs = wrappedJobArgs[JobArgs]{
		Args: args,
		Meta: JobMeta{
			TenantID: options.TenantID,
		},
	}

	// marhsal args to json
	if payload, err = json.Marshal(&wrappedArgs); err != nil {
		return nil, fmt.Errorf("json marhsal failed, %w", err)
	}

	var (
		task        = asynq.NewTask(args.Kind(), payload)
		taskOptions = []asynq.Option{
			applyAsynqOption(options.TaskID, asynq.TaskID),
			applyAsynqOption(options.Queue, asynq.Queue),
			applyAsynqOption(options.Timeout, asynq.Timeout),
			applyAsynqOption(options.MaxRetry, asynq.MaxRetry),
			applyAsynqOption(options.ProcessIn, asynq.ProcessIn),
			applyAsynqOption(options.ProcessAt, asynq.ProcessAt),
			applyAsynqOption(options.Deadline, asynq.Deadline),
			applyAsynqOption(options.Retention, asynq.Retention),
			applyAsynqOption(options.Group, asynq.Group),
			applyAsynqOption(options.Unique, asynq.Unique),
		}
	)

	// remove empty options
	taskOptions = slices.DeleteFunc(taskOptions, func(v asynq.Option) bool {
		return v == nil
	})

	var taskInfo *asynq.TaskInfo

	// enqueue using asynq client
	if taskInfo, err = c.client.EnqueueContext(ctx, task, taskOptions...); err != nil {
		return nil, fmt.Errorf("enqueu task with name %s failed, %w", args.Kind(), err)
	}

	return &Result[JobArgs]{
		ID:            taskInfo.ID,
		Queue:         taskInfo.Queue,
		Args:          args,
		MaxRetry:      taskInfo.MaxRetry,
		Retried:       taskInfo.Retried,
		LastErr:       taskInfo.LastErr,
		LastFailedAt:  taskInfo.LastFailedAt,
		Timeout:       taskInfo.Timeout,
		Deadline:      taskInfo.Deadline,
		Group:         taskInfo.Group,
		NextProcessAt: taskInfo.NextProcessAt,
		IsOrphaned:    taskInfo.IsOrphaned,
		Retention:     taskInfo.Retention,
		CompletedAt:   taskInfo.CompletedAt,
		Result:        taskInfo.Result,
	}, nil
}

// Close implements Client interface
func (c *client) Close(_ context.Context) error {
	return c.client.Close()
}
