// Package worker (task_send_verify_email) provides functions to distribute
// and process verify email tasks.
package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/utils"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	// TaskSendVerifyEmail represents the name of the task that sends the email verification email.
	TaskSendVerifyEmail = "task:send_verify_email"
)

// PayloadSendVerifyEmail provides the userID.
type PayloadSendVerifyEmail struct {
	UserID string `json:"user_id"`
}

// DistributeTaskSendVerifyEmail enqueues the given task to be processed by a worker. It returns an error if the task could
// not be enqueued.
func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("queue", info.Queue).
		Int("max_retry", info.MaxRetry).
		Msg("enqueued task")

	return nil
}

// ProcessTaskSendVerifyEmail processes a 'TaskSendVerifyEmail' task.
func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(
	ctx context.Context,
	task *asynq.Task,
) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUser(ctx, payload.UserID)
	if err != nil {
		return fmt.Errorf("failed to get error: %w", err)
	}

	now := time.Now()
	verifyEmail, err := processor.store.CreateVerifyEmail(ctx, &models.VerifyEmail{
		UserID:     user.ID,
		Email:      user.Contact.Email,
		SecretCode: utils.RandomString(64), // TODO: Substitute value with a token-based string
		CreatedAt:  now,
		ExpiredAt:  now.Add(15 * time.Minute),
	})
	if err != nil {
		return fmt.Errorf("failed to create verify email: %w", err)
	}

	verifyURL := fmt.Sprintf("http://localhost:8080/v1/verify_email?email_id=%d&secret_code=%s", verifyEmail.ID, verifyEmail.SecretCode)
	subject := "Welcome to Andela"
	content := fmt.Sprintf(`Hello %s, <br/>
	Thank you for registering with us! <br/>
	Please <a href="%s">Click here</a> to verify your email address.<br/>
	`, user.ID, verifyURL)
	to := []string{verifyEmail.Email}
	err = processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}

	log.Info().Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("email", user.Contact.Email).
		Msg("processed task")

	return nil
}
