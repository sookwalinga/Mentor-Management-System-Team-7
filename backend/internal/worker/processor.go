// Package worker (processor) provides a task processor for handling background tasks using the Asynq task queue.
package worker

import (
	"context"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/mail"
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	// QueueCritical is the name of the critical queue.
	QueueCritical = "critical"

	// QueueDefault is the name of the default queue.
	QueueDefault = "default"
)

// TaskProcessor is an interface for a worker that processes tasks.
type TaskProcessor interface {
	// Start starts the RedisTaskProcessor.
	Start() error

	// ProcessTaskSendResetPasswordEmail processes a 'TaskSendResetPasswordEmail' task.
	ProcessTaskSendResetPasswordEmail(
		ctx context.Context,
		task *asynq.Task,
	) error
}

// RedisTaskProcessor is a struct representing the task processor that implements the TaskProcessor interface.
// It is responsible for starting the Asynq task queue and processing tasks that are sent to it.
type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
	mailer mail.EmailSender
}

// NewRedisTaskProcessor creates a new RedisTaskProcessor.
// It initializes a new Asynq server with Redis client options, queue configurations, error handler and logger.
func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store, mailer mail.EmailSender) TaskProcessor {
	logger := NewLogger()
	redis.SetLogger(logger)

	server := asynq.NewServer(redisOpt, asynq.Config{
		Queues: map[string]int{
			QueueCritical: 10,
			QueueDefault:  5,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Error().Err(err).
				Str("type", task.Type()).
				Bytes("payload", task.Payload()).
				Msg("process task failed")
		}),
		Logger: logger,
	})
	return &RedisTaskProcessor{
		server: server,
		store:  store,
		mailer: mailer,
	}
}

// Start starts the RedisTaskProcessor.
func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendResetPasswordEmail, processor.ProcessTaskSendResetPasswordEmail)

	return processor.server.Start(mux)
}
