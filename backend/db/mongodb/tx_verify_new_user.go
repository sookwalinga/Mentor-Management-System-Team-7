// Package mongodb.tx_verify_new_user defines the database transaction
// to create a user and queue verify email transaction.
package mongodb

import (
	"context"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUserTx encapsulates the db transaction to create a user and
// queue a send verification email to new user.
func (mc *MongoClient) CreateUserTx(
	ctx context.Context,
	user *models.User,
	afterCreate func(*models.User) error,
) error {
	var result *models.User

	_, err := mc.execTx(ctx, func(sessionCtx mongo.SessionContext) (interface{}, error) {
		var err error

		result, err = mc.CreateUser(ctx, user)
		if err != nil {
			return nil, err
		}

		err = afterCreate(result)
		if err != nil {
			return nil, err
		}

		return result, nil
	})

	return err
}
