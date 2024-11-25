package middlewares

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLoggingMiddleware(logger *zap.Logger) asynq.MiddlewareFunc {
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			var (
				taskID, _     = asynq.GetTaskID(ctx)
				queueName, _  = asynq.GetQueueName(ctx)
				retryCount, _ = asynq.GetRetryCount(ctx)
				retryMax, _   = asynq.GetMaxRetry(ctx)
			)

			var (
				logger = logger.With(
					zap.String("task", taskID),
					zap.String("queue", queueName),
					zap.Int("retry_count", retryCount),
					zap.Int("retry_max", retryMax),
					zap.ByteString("payload", t.Payload()),
					zap.String("type", t.Type()),
				)
				start = time.Now()
				err   error
			)

			logger.Info("start processing")

			defer func() {
				var (
					logFields = []zapcore.Field{
						zap.Duration("elapsed", time.Since(start)),
					}
				)

				if err != nil {
					logger.Error("failed processing", append(logFields, zap.Error(err))...)
				} else {
					logger.Info("finished processing", logFields...)
				}
			}()

			if err = h.ProcessTask(ctx, t); err != nil {
				return err
			}
			return nil
		})
	}
}
