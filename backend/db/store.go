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

	// GetUserBy ID retrieves a user document from the collection by ID.
	GetUserByID(ctx context.Context, id string) (*models.User, error)

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

	// CreateFAQ inserts a new record into the faq collection.
	CreateFAQ(ctx context.Context, faq *models.Faq) (*models.Faq, error)

	// GetFAQ return a record by ID from the faq collection.
	GetFAQ(ctx context.Context, id string) (*models.Faq, error)

	// GetAllFAQs returns all records in faq collection.
	GetAllFAQs(ctx context.Context) ([]*models.Faq, error)

	// CreateDiscussion adds a new discussion document to the collection.
	CreateDiscussion(ctx context.Context, discussion *models.Discussion) (*models.Discussion, error)

	// GetDiscussion retrieves a discussion by its ID.
	GetDiscussion(ctx context.Context, discussionID string) (*models.Discussion, error)

	// ListDiscussions retrieves a list of discussions belonging to a particular owner with pagination.
	ListDiscussions(ctx context.Context, ownerID string, page int64, limit int64) ([]*models.Discussion, error)

	// UpdateDiscussion updates an existing discussion document in the collection.
	UpdateDiscussion(ctx context.Context, discussionID string, data map[string]interface{}) (*models.Discussion, error)

	// AddComment adds a new comment document to a discussion document in the collection.
	AddComment(ctx context.Context, discussionID string, comment *models.Comment) ([]models.Comment, error)
}
