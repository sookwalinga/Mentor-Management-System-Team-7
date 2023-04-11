package db

import (
	"context"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	// CreateUser adds a new user document to the collection.
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)

	// GetUserByID retrieves a user document from the collection by ID.
	GetUser(ctx context.Context, id string) (*models.User, error)

	// UpdateUserByID updates a user document in the collection by ID.
	UpdateUserByID(ctx context.Context, userID string, updateData map[string]interface{}) (*models.User, error)

	// DeleteUserByID removes a user document from the collection by ID.
	DeleteUserByID(ctx context.Context, id string) (*mongo.DeleteResult, error)

	// CreateVerifyEmail inserts a new record into the verify_emails collection.
	CreateVerifyEmail(ctx context.Context, verifyEmail *models.VerifyEmail) (*models.VerifyEmail, error)

	// GetVerifyEmailByID return a record by ID from the verify_emails collection.
	GetVerifyEmailByID(ctx context.Context, id string) (*models.VerifyEmail, error)

	// UpdateVerifyEmail updates and returns a record from the verify_emails collection
	UpdateVerifyEmail(ctx context.Context, id string, updateData map[string]interface{}) (*models.VerifyEmail, error)
}
