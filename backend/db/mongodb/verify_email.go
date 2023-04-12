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

// CreateVerifyEmail inserts a new record into the verify_emails collection.
func (mc *MongoClient) CreateVerifyEmail(ctx context.Context, verifyEmail *models.VerifyEmail) (*models.VerifyEmail, error) {
	result, err := mc.client.Database(DBName).Collection(VerifyEmailsCollection).InsertOne(ctx, verifyEmail)
	if err != nil {
		return nil, err
	}
	verifyEmailID := result.InsertedID.(string)
	return mc.GetVerifyEmail(ctx, verifyEmailID)
}

// GetVerifyEmail return a record by ID from the verify_emails collection.
func (mc *MongoClient) GetVerifyEmail(ctx context.Context, id string) (*models.VerifyEmail, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}

	verifyEmail := &models.VerifyEmail{}
	err = mc.client.Database(DBName).Collection(VerifyEmailsCollection).FindOne(ctx, filter).Decode(verifyEmail)
	return verifyEmail, err
}

// UpdateVerifyEmail updates and returns a record from the verify_emails collection
func (mc MongoClient) UpdateVerifyEmail(ctx context.Context, id string, updateData map[string]interface{}) (*models.VerifyEmail, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}

	updateDoc, err := bson.Marshal(updateData)
	if err != nil {
		return nil, err
	}

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	update := bson.M{"$set": updateDoc}

	updatedVerifyEmail := &models.VerifyEmail{}
	err = mc.client.Database(DBName).Collection(VerifyEmailsCollection).FindOneAndUpdate(ctx, filter, update, options).Decode(updatedVerifyEmail)
	return updatedVerifyEmail, err
}
