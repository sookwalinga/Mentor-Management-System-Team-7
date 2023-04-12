package token

import (
	"testing"
	"time"

	"github.com/ALCOpenSource/Mentor-Management-System-Team-7/backend/internal/utils"
	"github.com/stretchr/testify/require"
)

// TestPasetoMaker tests creating of paseto tokens.
func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	userID := utils.RandomUserID()
	userRole := utils.UserRole("Admin")
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, payload, err := maker.CreateToken(userID, userRole, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, userID, payload.UserID)
	require.Equal(t, userRole, payload.UserRole)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

// TestExpiredPasetoToken tests expiry of paseto tokens.
func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(
		utils.RandomUserID(), 
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
