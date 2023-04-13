// Package worker (task_send_reset_password_email) provides functions to distribute
// and process reset password email tasks.
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
	// TaskSendResetPasswordEmail represents the name of the task that sends the reset password email.
	TaskSendResetPasswordEmail = "task:send_reset_password_email"
)

// PayloadResetPasswordEmail provides the userEmail.
type PayloadResetPasswordEmail struct {
	UserID string `json:"user_id"`
}

// DistributeTaskSendResetPasswordEmail enqueues the given task to be processed by a worker. It returns an error if the task could
// not be enqueued.
func (distributor *RedisTaskDistributor) DistributeTaskSendResetPasswordEmail(
	ctx context.Context,
	payload *PayloadResetPasswordEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSendResetPasswordEmail, jsonPayload, opts...)
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

// ProcessTaskSendResetPasswordEmail processes a 'TaskSendResetPasswordEmail' task.
func (processor *RedisTaskProcessor) ProcessTaskSendResetPasswordEmail(
	ctx context.Context,
	task *asynq.Task,
) error {
	var payload PayloadResetPasswordEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUser(ctx, payload.UserID)
	if err != nil {
		return fmt.Errorf("failed to get error: %w", err)
	}

	now := time.Now()
	resetPassword, err := processor.store.CreateUserAction(ctx, &models.UserAction{
		UserID:     user.ID,
		Email:      user.Contact.Email,
		SecretCode: utils.RandomString(64), // TODO: Substitute value with a token-based string
		ActionType: "reset_password",
		CreatedAt:  now,
		ExpiredAt:  now.Add(15 * time.Minute),
	})
	if err != nil {
		return fmt.Errorf("failed to create reset password email: %w", err)
	}

	resetPasswordURL := fmt.Sprintf("http://localhost:8080/v1/reset_password?email_id=%d&secret_code=%s", resetPassword.ID, resetPassword.SecretCode)
	subject := "Request to reset password"
	content := fmt.Sprintf(`Hello %s, <br/>
	Your request to change password has been processed! <br/>
	Please <a href="%s">Click here</a> to reset your password.<br/>
	`, user.ID, resetPasswordURL)
	to := []string{resetPassword.Email}
	err = processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send reset password email: %w", err)
	}

	log.Info().Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("email", user.Contact.Email).
		Msg("processed task")

	return nil
}
