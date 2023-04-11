package mongodb

import (
	"context"
	"fmt"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DBName                 = ""
	UsersCollection        = "users"
	VerifyEmailsCollection = "verify_emails"
)

// MongoClient defines a Mongodb-based client.
type MongoClient struct {
	client *mongo.Client
}

// New instantiates a new Mongodb-based client.
func NewMongoClient(client *mongo.Client) db.Store {
	return &MongoClient{client}
}

// execTx prepares db transaction for execution
func (mc *MongoClient) execTx(ctx context.Context, txFunc func(sessionCtx mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	// Create a new session for the transaction
	session, err := mc.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	// Start the transaction
	result, err := session.WithTransaction(context.Background(), txFunc)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Transaction completed: %v\n", result)
	return result, nil
}
