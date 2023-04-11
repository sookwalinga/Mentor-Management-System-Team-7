package mongodb

import (
	"context"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"go.mongodb.org/mongo-driver/mongo"
)

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
