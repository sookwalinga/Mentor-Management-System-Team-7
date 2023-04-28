// Package mongodb (mongo) contains the wrapper for the Mongo Client
// implementing the Store interface.
package mongodb

import (
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// DBName defines database name.
	DBName = "MMS-Team7"

	// UsersCollection defines users collection name.
	UsersCollection = "users"

	// UserActionsCollection defines user_actions collection name.
	UserActionsCollection = "user_actions"

	// FAQCollection defines FAQ collection name.
	FAQCollection = "faq"

	// DiscussionsCollection defines Discussion collection name.
	DiscussionsCollection = "discussion_forum"
)

// MongoClient defines a Mongodb-based client.
type MongoClient struct {
	client *mongo.Client
}

// NewMongoClient instantiates a new Mongodb-based client.
func NewMongoClient(client *mongo.Client) db.Store {
	return &MongoClient{client}
}
