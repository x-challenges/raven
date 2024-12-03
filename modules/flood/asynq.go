package flood

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

// ApplyAsynqOption
func applyAsynqOption[T any](v *T, fn func(v T) asynq.Option) asynq.Option {
	if v != nil {
		return fn(*v)
	}
	return nil
}

// NewAsynqTask
func newAsynqTask[T JobArgs](_ Worker[T], payload []byte) *asynq.Task {
	var args T

	return asynq.NewTask(args.Kind(), payload, nil)
}

// AsynqJobMetaExtractor
func asynqJobMetaExtractor[T any](ctx context.Context, src func(ctx context.Context) (T, bool)) T {
	res, _ := src(ctx)
	return res
}

// NewAsynqHandlerFunc
func newAsynqHandlerFunc[T JobArgs](worker Worker[T]) asynq.HandlerFunc {
	return func(ctx context.Context, t *asynq.Task) error {
		var (
			args        T
			wrappedArgs wrappedJobArgs[T]
			err         error
		)

		// unmarshal asynq task payload to job args
		if err = json.Unmarshal(t.Payload(), &wrappedArgs); err != nil {
			return err
		}

		args = wrappedArgs.Args

		var job = &Job[T]{
			Args: args,
			Meta: JobMeta{
				TenantID: wrappedArgs.Meta.TenantID,
				TaskID:   asynqJobMetaExtractor(ctx, asynq.GetTaskID),
				MaxRetry: asynqJobMetaExtractor(ctx, asynq.GetMaxRetry),
				Attempt:  asynqJobMetaExtractor(ctx, asynq.GetRetryCount),
				Queue:    asynqJobMetaExtractor(ctx, asynq.GetQueueName),
			},
		}

		return worker.Work(ctx, job)
	}
}

// NewAsynqErrorHandlerFunc
func newAsynqErrorHandlerFunc[T JobArgs](logger *zap.Logger, worker Worker[T]) asynq.ErrorHandlerFunc {
	return func(ctx context.Context, t *asynq.Task, err error) {
		var (
			args        T
			wrappedArgs wrappedJobArgs[T]
		)

		// unmarshal asynq task payload to job args
		if err = json.Unmarshal(t.Payload(), &wrappedArgs); err != nil {
			logger.Error("unmarshal json faield, %v", zap.Error(err))
			return
		}

		args = wrappedArgs.Args

		var job = &Job[T]{
			Args: args,
			Meta: JobMeta{
				TenantID: wrappedArgs.Meta.TenantID,
				TaskID:   asynqJobMetaExtractor(ctx, asynq.GetTaskID),
				MaxRetry: asynqJobMetaExtractor(ctx, asynq.GetMaxRetry),
				Attempt:  asynqJobMetaExtractor(ctx, asynq.GetRetryCount),
				Queue:    asynqJobMetaExtractor(ctx, asynq.GetQueueName),
			},
		}

		// exec error handler for worker
		worker.ErrorHandler(ctx, job, err)
	}
}

// NewAsynqOptionFromWorker
func newAsynqOptionFromWorker[T JobArgs](_ Worker[T]) []asynq.Option {
	var (
		options = []asynq.Option{
			// TODO: apply options from worker ...
		}
	)

	return options
}
