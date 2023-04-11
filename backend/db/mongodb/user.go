package mongodb

import (
	"context"
	"errors"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser adds a new user document to the collection.
func (mc *MongoClient) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	result, err := mc.client.Database(DBName).Collection(UsersCollection).InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	userID := result.InsertedID.(string)
	return mc.GetUser(ctx, userID)
}

// GetUserByID retrieves a user document from the collection by ID.
func (mc *MongoClient) GetUser(ctx context.Context, id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	var user models.User
	err = mc.client.Database(DBName).Collection(UsersCollection).
		FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, db.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUserByID updates a user document in the collection by ID.
func (mc *MongoClient) UpdateUserByID(
	ctx context.Context,
	userID string,
	updateData map[string]interface{},
) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}

	// Convert the updateData map to a BSON document
	updateDoc, err := bson.Marshal(updateData)
	if err != nil {
		return nil, err
	}

	// Create an options document to specify which fields to update
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	update := bson.M{"$set": updateDoc}

	// Perform the update operation and store the updated user in the updatedUser variable
	updatedUser := &models.User{}
	err = mc.client.Database(DBName).Collection(UsersCollection).
		FindOneAndUpdate(ctx, filter, update, opts).Decode(updatedUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// DeleteUserByID removes a user document from the collection by ID.
func (mc *MongoClient) DeleteUserByID(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	return mc.client.Database(DBName).Collection(UsersCollection).DeleteOne(ctx, filter)
}
