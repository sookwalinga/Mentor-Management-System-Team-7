// Package mongodb (faq) defines functions to create and get FAQ records from DB.
package mongodb

import (
	"context"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateFAQ inserts a new record into the faq collection.
func (mc *MongoClient) CreateFAQ(ctx context.Context, faq *models.Faq) (*models.Faq, error) {
	result, err := mc.client.Database(DBName).Collection(FAQCollection).InsertOne(ctx, faq)
	if err != nil {
		return nil, err
	}
	userActionID := result.InsertedID.(primitive.ObjectID)
	return mc.GetFAQ(ctx, userActionID.Hex())
}

// GetFAQ return a record by ID from the faq collection.
func (mc *MongoClient) GetFAQ(ctx context.Context, id string) (*models.Faq, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}

	faq := &models.Faq{}
	err = mc.client.Database(DBName).Collection(FAQCollection).FindOne(ctx, filter).Decode(faq)
	return faq, err
}

// GetAllFAQs returns all records in faq collection.
func (mc *MongoClient) GetAllFAQs(ctx context.Context) ([]*models.Faq, error) {
	var faqs []*models.Faq
	cursor, err := mc.client.Database(DBName).Collection(FAQCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var faq models.Faq
		err := cursor.Decode(&faq)
		if err != nil {
			return nil, err
		}
		faqs = append(faqs, &faq)
	}

	return faqs, nil
}
