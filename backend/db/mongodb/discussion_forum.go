// Package mongodb (discussion_forum) includes db functions for the discussion forum.
package mongodb

import (
	"context"
	"fmt"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateDiscussion adds a new discussion document to the collection.
func (mc *MongoClient) CreateDiscussion(ctx context.Context, discussion *models.Discussion) (*models.Discussion, error) {
	result, err := mc.client.Database(DBName).Collection(DiscussionsCollection).InsertOne(ctx, discussion)
	if err != nil {
		return nil, err
	}
	discussionID := result.InsertedID.(primitive.ObjectID)
	return mc.GetDiscussion(ctx, discussionID.Hex())
}

// GetDiscussion retrieves a discussion by its ID.
func (mc *MongoClient) GetDiscussion(ctx context.Context, discussionID string) (*models.Discussion, error) {
	objectID, err := primitive.ObjectIDFromHex(discussionID)
	if err != nil {
		return nil, err
	}

	var discussion models.Discussion
	err = mc.client.Database(DBName).Collection(DiscussionsCollection).
		FindOne(ctx, bson.M{"_id": objectID}).
		Decode(&discussion)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, db.ErrRecordNotFound
		}
		return nil, err
	}

	return &discussion, nil
}

// ListDiscussions retrieves a list of discussions belonging to a particular owner with pagination.
func (mc *MongoClient) ListDiscussions(ctx context.Context, ownerID string, page int64, limit int64) ([]*models.Discussion, error) {
	objectID, err := primitive.ObjectIDFromHex(ownerID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"owner_id": objectID}
	opts := options.Find().SetSkip((page - 1) * limit).SetLimit(limit)
	cursor, err := mc.client.Database(DBName).Collection(DiscussionsCollection).Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var discussions []*models.Discussion
	for cursor.Next(ctx) {
		var discussion models.Discussion
		if err := cursor.Decode(&discussion); err != nil {
			return nil, err
		}
		discussions = append(discussions, &discussion)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return discussions, nil
}

// UpdateDiscussion updates an existing discussion document in the collection.
func (mc *MongoClient) UpdateDiscussion(ctx context.Context, discussionID string, data map[string]interface{}) (*models.Discussion, error) {
	objectID, err := primitive.ObjectIDFromHex(discussionID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": data}
	updatedDiscussion := &models.Discussion{}

	// Create an options document to specify which fields to update
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = mc.client.Database(DBName).Collection(DiscussionsCollection).FindOneAndUpdate(ctx, filter, update, opts).Decode(updatedDiscussion)
	return updatedDiscussion, err
}

// AddComment adds a new comment document to a discussion document in the collection.
func (mc *MongoClient) AddComment(ctx context.Context, discussionID string, comment *models.Comment) ([]models.Comment, error) {
	objectID, err := primitive.ObjectIDFromHex(discussionID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$push": bson.M{"comments": comment}}
	updatedDiscussion := &models.Discussion{}

	// Create an options document to specify which fields to update
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetProjection(bson.M{"_id": 0, "comments": 1})
	err = mc.client.Database(DBName).Collection(DiscussionsCollection).FindOneAndUpdate(ctx, filter, update, opts).Decode(updatedDiscussion)

	fmt.Println(updatedDiscussion)

	return updatedDiscussion.Comments, err
}
