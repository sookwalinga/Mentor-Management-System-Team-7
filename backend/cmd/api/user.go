// Package api (user) contains handlers for user data.
package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/token"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/utils"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/worker"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

type changeUserPasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required,min=8"`
	NewPassword     string `json:"new_password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_new_password" binding:"required,min=8,eqfield=NewPassword"`
}

type changeUserPasswordRequestID struct {
	ID string `uri:"id" binding:"required,min=1"`
}

func (server *Server) changeUserPassword(ctx *gin.Context) {
	var reqID changeUserPasswordRequestID
	if err := bindJSONWithValidation(ctx, ctx.ShouldBindUri(&reqID), validator.New()); err != nil {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if reqID.ID != authPayload.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse("mismatched user"))
		return
	}

	var req changeUserPasswordRequest
	if err := bindJSONWithValidation(ctx, ctx.ShouldBindJSON(&req), validator.New()); err != nil {
		return
	}

	user, err := server.store.GetUser(ctx, authPayload.UserID)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrRecordNotFound):
			ctx.JSON(http.StatusNotFound, errorResponse(db.ErrRecordNotFound.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse("failed to fetch user profile"))
		}
		return
	}

	err = utils.CheckPassword(req.CurrentPassword, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse("invalid login credentials"))
		return
	}

	hashedPassword, err := utils.HashedPassword(req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to hash password"))
		return
	}

	updateUserParams := map[string]interface{}{
		"hashed_password":     hashedPassword,
		"password_changed_at": time.Now(),
	}

	_, err = server.store.UpdateUser(ctx, authPayload.UserID, updateUserParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to update user's password"))
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
	if err := bindJSONWithValidation(ctx, ctx.ShouldBindJSON(&req), validator.New()); err != nil {
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrRecordNotFound):
			ctx.JSON(http.StatusNotFound, errorResponse(db.ErrRecordNotFound.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse("failed to fetch user profile"))
		}
		return
	}

	now := time.Now()
	resetPassword, err := server.store.CreateUserAction(ctx, &models.UserAction{
		UserID:     user.ID,
		Email:      user.Contact.Email,
		SecretCode: utils.RandomString(64), // TODO: Substitute value with a token-based string
		ActionType: "reset_password",
		CreatedAt:  now,
		ExpiredAt:  now.Add(15 * time.Minute),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to create user action"))
		return
	}

	task := &worker.PayloadResetPasswordEmail{
		ID:        resetPassword.ID.Hex(),
		UserID:    user.ID.Hex(),
		UserEmail: user.Contact.Email,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(5 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	err = server.taskDistributor.DistributeTaskSendResetPasswordEmail(ctx, task, opts...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to enqueue task"))
		return
	}

	ctx.JSON(http.StatusOK, envelop{"result": "reset password email sent"})

	log.Info().
		Str("user_id", user.ID.Hex()).
		Str("ip_address", ctx.ClientIP()).
		Str("user_agent", ctx.Request.UserAgent()).
		Str("request_method", ctx.Request.Method).
		Str("request_path", ctx.Request.URL.Path).
		Msg("task 'reset password' enqueued")
}

// Login
type userLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (server *Server) login(ctx *gin.Context) {
	var req userLogin
	// Validate request.
	if err := bindJSONWithValidation(ctx, ctx.ShouldBindJSON(&req), validator.New()); err != nil {
		return
	}
	// Get user by email.
	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		switch {
		case errors.Is(err, db.ErrRecordNotFound):
			ctx.JSON(http.StatusNotFound, errorResponse(db.ErrRecordNotFound.Error()))
		default:
			ctx.JSON(http.StatusInternalServerError, errorResponse("failed to fetch user profile"))
		}
		return
	}
	// Check password.
	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse("invalid login credentials"))
		return
	}
	// Create token.
	token, payload, err := server.tokenMaker.CreateToken(user.ID.Hex(), user.Role, 24*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to create access token"))
		return
	}

	// Return token.
	ctx.JSON(http.StatusOK,
		envelop{
			"data": gin.H{
				"data":    user,
				"token":   token,
				"payload": payload,
			},
		},
	)
	log.Info().
		Str("user_id", user.ID.Hex()).
		Str("ip_address", ctx.ClientIP()).
		Str("user_agent", ctx.Request.UserAgent()).
		Str("request_method", ctx.Request.Method).
		Str("request_path", ctx.Request.URL.Path).
		Msg("user logged in")
}

// Define the Google Sign-in route handler
func (server *Server) googleLogin(w http.ResponseWriter, r *http.Request) {
	url := server.googleConfig.AuthCodeURL(server.config.GoogleRandomString, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

type profile struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	EmailVerified bool   `json:"email_verified"`
}

// Define the Google Sign-in callback route handler
func (server *Server) googleLoginCallback(ctx *gin.Context) {
	// Check state is valid.
	state := ctx.Query("state")
	if state != server.config.GoogleRandomString {
		ctx.JSON(http.StatusInternalServerError, errorResponse("invalid state value"))
		return
	}

	// Exchange the authorization code for an access token and ID token
	code := ctx.Query("code")
	token, err := server.googleConfig.Exchange(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to exchange code"))
		return
	}

	// Get the user's Google profile using the access token
	client := server.googleConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to get user info"))
		return
	}
	defer resp.Body.Close()

	// Parse the user's profile JSON
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to read response body"))
		return
	}
	userProfile := &profile{}
	if err := json.Unmarshal(body, userProfile); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to parse user profle"))
		return
	}

	// Retrieve user by email
	user, err := server.store.GetUserByEmail(ctx, userProfile.Email)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(db.ErrRecordNotFound.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to fetch user profile"))
		return
	}

	pasetoToken, payload, err := server.tokenMaker.CreateToken(user.ID.Hex(), user.Role, 24*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to create access token"))
		return
	}

	ctx.JSON(http.StatusOK, envelop{
		"data": gin.H{
			"user":         user,
			"payload":      payload,
			"access_token": pasetoToken,
		},
	})

	log.Info().
		Str("user_id", user.ID.Hex()).
		Str("ip_address", ctx.ClientIP()).
		Str("user_agent", ctx.Request.UserAgent()).
		Str("request_method", ctx.Request.Method).
		Str("request_path", ctx.Request.URL.Path).
		Msg("logged in user using google sigin-in")
}

type updateUserRequest struct {
	FirstName       *string `json:"first_name" binding:"min=2"`
	LastName        *string `json:"last_name" binding:"min=2"`
	About           *string `json:"about"`
	Website         *string `json:"website"`
	ProfileImageURL *string `json:"profile_image_url"`
	Country         *string `json:"country"`
	City            *string `json:"city"`
	GitHubURL       *string `json:"github_url"`
	LinkedInURL     *string `json:"linkedin_url"`
	TwitterURL      *string `json:"twitter_url"`
	InstagramURL    *string `json:"instagram_url"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := bindJSONWithValidation(ctx, ctx.ShouldBindJSON(&req), validator.New()); err != nil {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := map[string]interface{}{}

	if req.FirstName != nil {
		arg["first_name"] = *req.FirstName
	}
	if req.LastName != nil {
		arg["last_name"] = *req.LastName
	}
	if req.About != nil {
		arg["about"] = *req.About
	}
	if req.Website != nil {
		arg["contact.website"] = *req.Website
	}
	if req.ProfileImageURL != nil {
		arg["profile_image_url"] = *req.ProfileImageURL
	}
	if req.Country != nil {
		arg["contact.country"] = *req.Country
	}
	if req.City != nil {
		arg["contact.city"] = *req.City
	}
	if req.GitHubURL != nil {
		arg["socials.github_url.value"] = *req.GitHubURL
	}
	if req.LinkedInURL != nil {
		arg["socials.linkedin_url.value"] = *req.LinkedInURL
	}
	if req.TwitterURL != nil {
		arg["socials.twitter_url.value"] = *req.TwitterURL
	}
	if req.InstagramURL != nil {
		arg["socials.instagram_url.value"] = *req.InstagramURL
	}

	resp, err := server.store.UpdateUser(ctx, authPayload.UserID, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to update user's profile"))
		return
	}

	ctx.JSON(http.StatusOK, envelop{
		"data": gin.H{
			"user": resp,
		},
	})

	log.Info().
		Str("user_id", authPayload.UserID).
		Str("ip_address", ctx.ClientIP()).
		Str("user_agent", ctx.Request.UserAgent()).
		Str("request_method", ctx.Request.Method).
		Str("request_path", ctx.Request.URL.Path).
		Msg("updated user successfully")
}

func (server *Server) logout(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	expiredAt := authPayload.ExpiredAt
	duration := time.Until(expiredAt)

	err := server.cache.BlacklistSession(ctx, authPayload.ID.String(), duration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to blacklist access token"))
		return
	}

	ctx.JSON(http.StatusOK, envelop{"result": "Logged out user successfully"})
	log.Info().
		Str("user_id", authPayload.UserID).
		Str("ip_address", ctx.ClientIP()).
		Str("user_agent", ctx.Request.UserAgent()).
		Str("request_method", ctx.Request.Method).
		Str("request_path", ctx.Request.URL.Path).
		Msg("logged out user")
}
