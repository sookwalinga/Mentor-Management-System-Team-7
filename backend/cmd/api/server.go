package api

import (
	"fmt"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/token"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/utils"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/worker"

	"github.com/gin-gonic/gin"
)

// A Server serves HTTP requests for the banking system
type Server struct {
	config     utils.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer create a new HTTP server and setup routing.
func NewServer(config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		taskDistributor: taskDistributor,
	}

	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()

	// TODO: setup routes and its equivalent handlers
	// ex. router.POST("/users/login", s.loginUser)

	s.router = router
}

// Start run the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
