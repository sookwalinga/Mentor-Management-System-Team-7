// Package token (payload) defines the token's payload.
package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	// ErrInvalidToken is returned for invalid token
	ErrInvalidToken = errors.New("token is invalid")

	// ErrExpiredToken is returned for expired token
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token.
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	UserRole  string    `json:"user_role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload creates a new token payload with a specific username and duration.
func NewPayload(userID, userRole string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		UserRole:  userRole,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// Valid checks for the expiry-validity of the token.
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
