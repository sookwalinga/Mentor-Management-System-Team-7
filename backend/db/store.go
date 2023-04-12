// Package db.store defines the Store interface that a datastore
// needs to implement/satisfy.
package db

import (
	"context"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// A Store provides all functions to execute db queries
// and transactions.
type Store interface {
	// CreateUser adds a new user document to the collection.
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)

	// GetUserByID retrieves a user document from the collection by ID.
	GetUser(ctx context.Context, id string) (*models.User, error)

	// UpdateUser updates a user document in the collection by ID.
	UpdateUser(ctx context.Context, userID string, updateData map[string]interface{}) (*models.User, error)

	// DeleteUser removes a user document from the collection by ID.
	DeleteUser(ctx context.Context, id string) (*mongo.DeleteResult, error)

	// CreateVerifyEmail inserts a new record into the verify_emails collection.
	CreateVerifyEmail(ctx context.Context, verifyEmail *models.VerifyEmail) (*models.VerifyEmail, error)

	// GetVerifyEmailByID return a record by ID from the verify_emails collection.
	GetVerifyEmail(ctx context.Context, id string) (*models.VerifyEmail, error)

	// UpdateVerifyEmail updates and returns a record from the verify_emails collection
	UpdateVerifyEmail(ctx context.Context, id string, updateData map[string]interface{}) (*models.VerifyEmail, error)
}
