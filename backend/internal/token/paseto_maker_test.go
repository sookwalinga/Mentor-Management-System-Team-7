package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/utils"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestPasetoMaker tests creating of paseto tokens.
func TestPasetoMaker(t *testing.T) {
	config, err := utils.LoadConfig("../..")
	require.NoError(t, err)

	maker, err := NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	// userID := utils.RandomUserID().Hex()
	userID, err := primitive.ObjectIDFromHex("643a7dedbc8c7b338e50bd0f")
	require.NoError(t, err)

	userRole := "SuperAdmin"
	duration := 15 * time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, payload, err := maker.CreateToken(userID.Hex(), userRole, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	fmt.Println(token)
	fmt.Printf("\n\n")

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, userID.Hex(), payload.UserID)
	require.Equal(t, userRole, payload.UserRole)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

// TestExpiredPasetoToken tests expiry of paseto tokens.
func TestExpiredPasetoToken(t *testing.T) {
	config, err := utils.LoadConfig("../..")
	require.NoError(t, err)

	maker, err := NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(
		utils.RandomUserID().Hex(),
		utils.UserRole("Admin"),
		-time.Minute,
	)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
