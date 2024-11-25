package flood

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
)

func applyAsynqOption[T any](v *T, fn func(v T) asynq.Option) asynq.Option {
	if v != nil {
		return fn(*v)
	}
	return nil
}

func newAsynqTask[T JobArgs](_ Worker[T], payload []byte) *asynq.Task {
	var args T
	return asynq.NewTask(args.Kind(), payload, nil)
}

func asynqJobMetaExtractor[T any](ctx context.Context, src func(ctx context.Context) (T, bool)) T {
	res, _ := src(ctx)
	return res
}

func newAsynqHandlerFn[T JobArgs](worker Worker[T]) asynq.HandlerFunc {
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

func newAsynqOptionFromWorker[T JobArgs](_ Worker[T]) []asynq.Option {
	var (
		options = []asynq.Option{
			// TODO: apply options from worker ...
		}
	)
	return options
}
