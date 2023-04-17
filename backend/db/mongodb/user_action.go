// Package mongodb (verify_email) defines db queries for the
// verify_email collection.
package mongodb

import (
	"context"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUserAction inserts a new record into the user_actions collection.
func (mc *MongoClient) CreateUserAction(ctx context.Context, userAction *models.UserAction) (*models.UserAction, error) {
	result, err := mc.client.Database(DBName).Collection(UserActionsCollection).InsertOne(ctx, userAction)
	if err != nil {
		return nil, err
	}
	userActionID := result.InsertedID.(primitive.ObjectID)
	return mc.GetUserAction(ctx, userActionID.Hex())
}

// GetUserAction return a record by ID from the user_actions collection.
func (mc *MongoClient) GetUserAction(ctx context.Context, id string) (*models.UserAction, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}

	userAction := &models.UserAction{}
	err = mc.client.Database(DBName).Collection(UserActionsCollection).FindOne(ctx, filter).Decode(userAction)
	return userAction, err
}

// UpdateUserAction updates and returns a record from the user_actions collection
func (mc MongoClient) UpdateUserAction(ctx context.Context, id string, updateData map[string]interface{}) (*models.UserAction, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	update := bson.M{"$set": updateData}

	updateduserAction := &models.UserAction{}
	err = mc.client.Database(DBName).Collection(UserActionsCollection).FindOneAndUpdate(ctx, filter, update, options).Decode(updateduserAction)
	return updateduserAction, err
}
