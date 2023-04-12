// Package utils (random) defines general utilities for codebase.
package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

// User role
const (
	ADMIN  = "Admin"
	MENTOR = "Mentor" 
	MENTEE = "Mentee"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates random integer between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomUserID generates a random user name.
func RandomUserID() string {
	userID := primitive.NewObjectID()
	return userID.String()
}

// UserRole returns a role if found.
func UserRole(role string) string {
	roles := []string{ADMIN, MENTOR, MENTEE}
	for _, r := range roles {
		if role == r {
			return role
		}
	}
	return roles[0]
}

// RandomEmail generates a random email.
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
