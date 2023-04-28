// Package api (middleware) defines all the middlewares for the server.
package api

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func (server *Server) authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if authorizationHeader == "" {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error()))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error()))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error()))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error()))
			return
		}

		exists, err := server.cache.IsSessionBlacklisted(ctx, payload.ID.String())
		if err != nil || exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse("invalid token"))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

// ResponseBodyWriter that wraps the original gin.ResponseWriter
type ResponseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write capture the response body as it's being written by the next handler
func (w ResponseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Create a new buffer to capture the response body
		buffer := bytes.NewBufferString("")
		writer := ResponseBodyWriter{c.Writer, buffer}
		c.Writer = writer

		// Call the next handler
		c.Next()

		duration := time.Since(startTime)

		logger := log.Info()
		if c.Writer.Status() < http.StatusOK || c.Writer.Status() >= http.StatusBadRequest {
			body := buffer.Bytes()
			logger = log.Error().Bytes("body", body)
		} else if c.Writer.Status() >= http.StatusMultipleChoices && c.Writer.Status() < http.StatusBadRequest {
			body := buffer.Bytes()
			logger = log.Warn().Bytes("body", body)
		}

		logger.Str("protocol", "HTTP").
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status_code", c.Writer.Status()).
			Str("status_text", http.StatusText(c.Writer.Status())).
			Dur("duration", duration).
			Msg("received an HTTP request")
	}
}
