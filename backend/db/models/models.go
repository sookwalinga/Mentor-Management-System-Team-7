// Package models contains the data models used in the application.
// It defines the data structure for a user's profile and settings, including contact details, social media URLs,
// notification settings, support information, and email verification details.
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user's profile & settings.
type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	FullName          string             `bson:"full_name,omitempty"`
	Role              string             `bson:"role,omitempty"`
	About             string             `bson:"about,omitempty"`
	Contact           Contact            `bson:"contact,omitempty"`
	CreatedAt         time.Time          `bson:"created_at,omitempty"`
	ProfileImageURL   string             `bson:"profile_image_url,omitempty"`
	PasswordChangedAt time.Time          `bson:"password_changed_at,omitempty"`
	HashedPassword    string             `bson:"hashed_password,omitempty"`
	Socials           Socials            `bson:"socials,omitempty"`
	IsEmailVerified   bool               `bson:"is_email_verified,omitempty"`
}

// Contact represents user's contact details.
type Contact struct {
	Email    string `bson:"email,omitempty"`
	Website  string `bson:"website,omitempty"`
	Location string `bson:"location,omitempty"`
}

// SocialMediaURL represents the value and settings for socials.
type SocialMediaURL struct {
	Value     string `bson:"value,omitempty"`
	IsVisible bool   `bson:"is_visible,omitempty"`
}

// Socials represents user's social details.
type Socials struct {
	GitHubURL    SocialMediaURL `bson:"github_url,omitempty"`
	LinkedInURL  SocialMediaURL `bson:"linkedin_url,omitempty"`
	TwitterURL   SocialMediaURL `bson:"twitter_url,omitempty"`
	InstagramURL SocialMediaURL `bson:"instagram_url,omitempty"`
}

// A NotificationType represents notification settings.
type NotificationType struct {
	EmailEnabled bool `bson:"enable_email,omitempty"`
	InAppEnabled bool `bson:"enable_in_app,omitempty"`
}

// A Notifications represents user's notification details.
type Notifications struct {
	ID                           primitive.ObjectID `bson:"_id,omitempty"`
	AllNotifications             NotificationType   `bson:"all_notifications,omitempty"`
	ProgramNotifications         NotificationType   `bson:"program_notifications,omitempty"`
	TaskNotifications            NotificationType   `bson:"task_notifications,omitempty"`
	ApprovalRequestNotifications NotificationType   `bson:"approval_request_notifications,omitempty"`
	ReportsNotifications         NotificationType   `bson:"reports_notifications,omitempty"`
	PostNotifications            NotificationType   `bson:"post_notifications,omitempty"`
	CommentsNotifications        NotificationType   `bson:"comments_notifications,omitempty"`
	MentionsNotifications        NotificationType   `bson:"mentions_notifications,omitempty"`
	DirectMessageNotifications   NotificationType   `bson:"direct_message_notifications,omitempty"`
	CommentsOnPostNotifications  NotificationType   `bson:"comments_on_post_notifications,omitempty"`
}

// A Support represents support information.
type Support struct {
	Name  string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
	Title string `bson:"title,omitempty"`
	Body  string `bson:"body,omitempty"`
}

// UserAction represents user action details
type UserAction struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty"`
	Email      string             `bson:"email,omitempty"`
	SecretCode string             `bson:"secret_code,omitempty"`
	ActionType string             `bson:"action_type,omitempty"`
	IsUsed     bool               `bson:"is_used,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty"`
	ExpiredAt  time.Time          `bson:"expired_at,omitempty"`
}

// Faq represents faq details
type Faq struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Question string             `bson:"question" json:"question"`
	Answer   string             `bson:"answer" json:"answer"`
	Category string             `bson:"category" json:"category"`
}
