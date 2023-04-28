// Package api (faq) defines handlers for creating and retrieving FAQs.
package api

import (
	"net/http"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type createFAQRequest struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
	Category string `json:"category" binding:"required"`
}

func (server *Server) createFAQ(ctx *gin.Context) {
	var req createFAQRequest
	if err := bindJSONWithValidation(ctx, ctx.ShouldBindJSON(&req), validator.New()); err != nil {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.UserRole != "Admin" {
		ctx.JSON(http.StatusUnauthorized, errorResponse("not authorised to create faq"))
		return
	}

	faq := &models.Faq{
		Question: req.Question,
		Answer:   req.Answer,
		Category: req.Category,
	}

	resp, err := server.store.CreateFAQ(ctx, faq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to create FAQ"))
		return
	}

	ctx.JSON(http.StatusCreated, envelop{"data": resp})
	log.Info().
		Str("user_id", authPayload.UserID).
		Str("ip_address", ctx.ClientIP()).
		Str("user_agent", ctx.Request.UserAgent()).
		Str("request_method", ctx.Request.Method).
		Str("request_path", ctx.Request.URL.Path).
		Msg("FAQ created")
}

func (server *Server) getAllFAQs(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	resp, err := server.store.GetAllFAQs(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("failed to get all FAQs"))
		return
	}

	ctx.JSON(http.StatusOK, envelop{"data": resp})

	log.Info().
		Str("user_id", authPayload.UserID).
		Str("ip_address", ctx.ClientIP()).
		Str("user_agent", ctx.Request.UserAgent()).
		Str("request_method", ctx.Request.Method).
		Str("request_path", ctx.Request.URL.Path).
		Msg("FAQ created")
}
