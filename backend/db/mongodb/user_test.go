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

func createTestUser(t *testing.T) *models.User {
	// hashedPassword, err := utils.HashedPassword(utils.RandomString(10))
	hashedPassword, err := utils.HashedPassword("secretsz")
	require.NoError(t, err)

	user := models.User{
		ID:        utils.RandomUserID(),
		FirstName: utils.RandomString(4),
		LastName:  utils.RandomString(4),
		Role:      "Mentor",
		About:     utils.RandomString(12),
		Contact: models.Contact{
			Email:   utils.RandomEmail(),
			Website: utils.RandomString(5),
			Country: utils.RandomString(5),
			City:    utils.RandomString(5),
		},
		CreatedAt:         time.Now(),
		ProfileImageURL:   utils.RandomString(12),
		PasswordChangedAt: time.Now(),
		HashedPassword:    hashedPassword,
		IsEmailVerified:   false,
		Socials: models.Socials{
			GitHubURL: models.SocialMediaURL{
				Value:     fmt.Sprintf("https://github.com/%s", utils.RandomString(8)),
				IsVisible: true,
			},
			LinkedInURL: models.SocialMediaURL{
				Value:     fmt.Sprintf("https://www.linkedin.com/in/%s", utils.RandomString(8)),
				IsVisible: true,
			},
			TwitterURL: models.SocialMediaURL{
				Value:     fmt.Sprintf("https://twitter.com/%s", utils.RandomString(8)),
				IsVisible: true,
			},
			InstagramURL: models.SocialMediaURL{
				Value:     fmt.Sprintf("https://www.instagram.com/%s", utils.RandomString(8)),
				IsVisible: true,
			},
		},
	}

	userRecord, err := testStore.CreateUser(context.Background(), &user)
	require.NoError(t, err)
	require.NotEmpty(t, userRecord)
	compareUser(t, &user, userRecord)

	return userRecord
}

func TestCreateUser(t *testing.T) {
	createTestUser(t)
}

func compareUser(t *testing.T, user1, user2 *models.User) {
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Role, user2.Role)
	require.Equal(t, user1.About, user2.About)
	require.Equal(t, user1.Contact, user2.Contact)
	require.Equal(t, user1.Socials, user2.Socials)
	require.Equal(t, user1.ProfileImageURL, user2.ProfileImageURL)
	require.Equal(t, user1.IsEmailVerified, user2.IsEmailVerified)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestGetUser(t *testing.T) {
	user1 := createTestUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.ID.Hex())
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID.Hex(), user2.ID.Hex())
	compareUser(t, user1, user2)
}

func TestGetUserByEmail(t *testing.T) {
	user1 := createTestUser(t)
	user2, err := testStore.GetUserByEmail(context.Background(), user1.Contact.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID.Hex(), user2.ID.Hex())
	compareUser(t, user1, user2)
}

func TestUpdateUserOnlyFullName(t *testing.T) {
	oldUser := createTestUser(t)

	newFirstName := utils.RandomString(4)
	newLastName := utils.RandomString(4)
	updatedUser, err := testStore.UpdateUser(context.Background(), oldUser.ID.Hex(), map[string]interface{}{
		"first_name": newFirstName,
		"last_name":  newLastName,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FirstName, updatedUser.FirstName)
	require.NotEqual(t, oldUser.LastName, updatedUser.LastName)
	require.Equal(t, newFirstName, updatedUser.FirstName)
	require.Equal(t, newLastName, updatedUser.LastName)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createTestUser(t)

	newEmail := utils.RandomEmail()
	updatedUser, err := testStore.UpdateUser(context.Background(), oldUser.ID.Hex(), map[string]interface{}{
		"contact.email": newEmail,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Contact.Email, updatedUser.Contact.Email)
	require.Equal(t, newEmail, updatedUser.Contact.Email)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createTestUser(t)

	newPassword := utils.RandomString(6)
	newHashedPassword, err := utils.HashedPassword(newPassword)
	require.NoError(t, err)

	updatedUser, err := testStore.UpdateUser(context.Background(), oldUser.ID.Hex(), map[string]interface{}{
		"hashed_password": newHashedPassword,
	})

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
}
