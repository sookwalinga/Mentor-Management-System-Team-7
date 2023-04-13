// Package (user) contains handlers for user data.
package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/token"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/utils"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type changeUserPasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required,min=8"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_new_password" binding:"required,min=8,eqfield=NewPassword"`
}

func (server *Server) changeUserPassword(ctx *gin.Context) {
	var req changeUserPasswordRequest

	if err := BindJSONWithValidation(ctx, &req, validator.New()); err != nil {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := server.store.GetUser(ctx, authPayload.UserID)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrRecordNotFound):
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	err = utils.CheckPassword(req.CurrentPassword, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashedPassword(req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	updateUserParams := map[string]interface{}{
		"hashed_password":     hashedPassword,
		"password_changed_at": time.Now(),
	}

	_, err = server.store.UpdateUser(ctx, authPayload.UserID, updateUserParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": "password changed successfully"})
	log.Info().
		Str("user_id", user.ID.Hex()).
		Str("ip_address", ctx.ClientIP()).
		Str("user_agent", ctx.Request.UserAgent()).
		Str("request_method", ctx.Request.Method).
		Str("request_path", ctx.Request.URL.Path).
		Msg("password changed for user")
}

type forgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func (server *Server) forgotPassword(ctx *gin.Context) {
	var req forgotPasswordRequest

	if err := BindJSONWithValidation(ctx, &req, validator.New()); err != nil {
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrRecordNotFound):
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	task := &worker.PayloadResetPasswordEmail{
		UserID: user.ID.Hex(),
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(5 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	err = server.taskDistributor.DistributeTaskSendResetPasswordEmail(ctx, task, opts...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	log.Info().
		Str("user_id", user.ID.Hex()).
		Str("ip_address", ctx.ClientIP()).
		Str("user_agent", ctx.Request.UserAgent()).
		Str("request_method", ctx.Request.Method).
		Str("request_path", ctx.Request.URL.Path).
		Msg("task 'reset password' enqueued")
}
