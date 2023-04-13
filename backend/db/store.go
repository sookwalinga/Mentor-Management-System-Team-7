// Package db (store) defines the Store interface that a datastore
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

	// GetUserByEmail retrieves a user document from the collection by email.
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)

	// UpdateUser updates a user document in the collection by ID.
	UpdateUser(ctx context.Context, userID string, updateData map[string]interface{}) (*models.User, error)

	// DeleteUser removes a user document from the collection by ID.
	DeleteUser(ctx context.Context, id string) (*mongo.DeleteResult, error)

	// CreateUserAction inserts a new record into the user_actions collection.
	CreateUserAction(ctx context.Context, userAction *models.UserAction) (*models.UserAction, error)

	// GetUserAction return a record by ID from the user_actions collection.
	GetUserAction(ctx context.Context, id string) (*models.UserAction, error)

	// UpdateUserAction updates and returns a record from the user_actions collection
	UpdateUserAction(ctx context.Context, id string, updateData map[string]interface{}) (*models.UserAction, error)
}
