// Package api provides an HTTP server implementation for the Mentor-Management-System-Team-7
// backend application. It defines a Server struct that serves HTTP requests and sets up the routing.
// The package utilizes the following internal packages: db, token, utils, and worker.
package api

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/token"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/utils"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/worker"

	"github.com/gin-gonic/gin"
)

// A Server serves HTTP requests for the banking system
type Server struct {
	config          utils.Config
	store           db.Store
	router          *gin.Engine
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
	swaggerFiles fs.FS
}

// NewServer create a new HTTP server and setup routing.
func NewServer(
	config utils.Config, 
	store db.Store, 
	taskDistributor worker.TaskDistributor, 
	swaggerFiles fs.FS,
) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
		swaggerFiles: swaggerFiles,
	}

	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.Use(loggerMiddleware())
	router.POST("/api/v1/forgot_password", s.forgotPassword)
	router.POST("/api/v1/auth/login", s.login)

	fsysHandler := http.FileServer(http.FS(s.swaggerFiles))
	router.GET("/swagger/*any", gin.WrapH(http.StripPrefix("/swagger/", fsysHandler)))

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))
	authRoutes.PATCH("/api/v1/users/:id/change_password", s.changeUserPassword)
	authRoutes.POST("/api/v1/faqs", s.createFAQ)
	authRoutes.GET("/api/v1/faqs", s.getAllFAQs)
	authRoutes.POST("/api/v1/user/:id", s.updateUser)

	s.router = router
}

// Start run the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

type envelop map[string]interface{}
