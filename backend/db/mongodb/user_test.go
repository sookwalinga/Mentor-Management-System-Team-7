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
	hashedPassword, err := utils.HashedPassword("secrets")
	require.NoError(t, err)

	user := models.User{
		ID:       utils.RandomUserID(),
		FullName: utils.RandomString(8),
		Role:     "Admin",
		About:    utils.RandomString(12),
		Contact: models.Contact{
			Email:    utils.RandomEmail(),
			Website:  utils.RandomString(5),
			Location: utils.RandomString(5),
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

	require.Equal(t, user.FullName, userRecord.FullName)
	require.Equal(t, user.Role, userRecord.Role)
	require.Equal(t, user.About, userRecord.About)
	require.Equal(t, user.Contact, userRecord.Contact)
	require.WithinDuration(t, user.CreatedAt, userRecord.CreatedAt, time.Millisecond)
	require.Equal(t, user.ProfileImageURL, userRecord.ProfileImageURL)
	require.WithinDuration(t, user.PasswordChangedAt, userRecord.PasswordChangedAt, time.Millisecond)
	require.Equal(t, user.HashedPassword, userRecord.HashedPassword)
	require.Equal(t, user.Socials, userRecord.Socials)
	require.Equal(t, user.IsEmailVerified, userRecord.IsEmailVerified)

	return userRecord
}

func TestCreateUser(t *testing.T) {
	createTestUser(t)
}
