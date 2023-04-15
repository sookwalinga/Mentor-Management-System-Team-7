package mongodb

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/db/models"
	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/utils"
	"github.com/stretchr/testify/require"
)

func createTestUserAction(t *testing.T) *models.UserAction {
	user := createTestUser(t)

	now := time.Now()
	userAction := &models.UserAction{
		UserID:     user.ID,
		Email:      user.Contact.Email,
		SecretCode: utils.RandomString(32),
		ActionType: "verify_email",
		IsUsed:     false,
		CreatedAt:  now,
		ExpiredAt:  now.Add(15 * time.Minute),
	}

	actionRecord, err := testStore.CreateUserAction(context.Background(), userAction)
	require.NoError(t, err)
	require.NotEmpty(t, actionRecord)

	require.NotNil(t, actionRecord.ID.Hex())
	require.Equal(t, userAction.UserID, actionRecord.UserID)
	require.Equal(t, userAction.Email, actionRecord.Email)
	require.Equal(t, userAction.SecretCode, actionRecord.SecretCode)
	require.Equal(t, userAction.ActionType, actionRecord.ActionType)
	require.False(t, userAction.IsUsed)
	require.WithinDuration(t, userAction.CreatedAt, actionRecord.CreatedAt, time.Second)
	require.WithinDuration(t, userAction.ExpiredAt, actionRecord.ExpiredAt, time.Second)

	return actionRecord
}

func TestCreateUserAction(t *testing.T) {
	c := createTestUserAction(t)

	fmt.Println(c)
}

func TestUpdateUserAction(t *testing.T) {
	action := createTestUserAction(t)

	updatedAction, err := testStore.UpdateUserAction(context.Background(), action.ID.Hex(), map[string]interface{}{
		"is_used": true,
	})
	require.NoError(t, err)
	require.True(t, updatedAction.IsUsed)
}
