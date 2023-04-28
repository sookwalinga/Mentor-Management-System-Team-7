// Package worker (distributor) provides a task distributor for sending background tasks to redies
// using the Asynq task queue.
package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

// TaskDistributor defines the inteface required to
// distribute asynchronous tasks.
type TaskDistributor interface {
	DistributeTaskSendResetPasswordEmail(
		ctx context.Context,
		payload *PayloadResetPasswordEmail,
		opts ...asynq.Option,
	) error
}

// RedisTaskDistributor defines and wrap a asynq client
// to distribute task to redis.
type RedisTaskDistributor struct {
	client *asynq.Client
}

// NewRedisTaskDistributor instantiates a RedisTaskDistributor object.
func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
	}
}
