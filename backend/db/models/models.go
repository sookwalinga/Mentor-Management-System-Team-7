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
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FullName          string             `bson:"full_name,omitempty" json:"full_name,omitempty"`
	Role              string             `bson:"role,omitempty" json:"role,omitempty"`
	About             string             `bson:"about,omitempty" json:"about,omitempty"`
	Contact           Contact            `bson:"contact,omitempty" json:"contact,omitempty"`
	CreatedAt         time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	ProfileImageURL   string             `bson:"profile_image_url,omitempty" json:"profile_image_url,omitempty"`
	PasswordChangedAt time.Time          `bson:"password_changed_at,omitempty" json:"password_changed_at,omitempty"`
	HashedPassword    string             `bson:"hashed_password,omitempty" json:"-"`
	Socials           Socials            `bson:"socials,omitempty" json:"socials,omitempty"`
	IsEmailVerified   bool               `bson:"is_email_verified,omitempty" json:"is_email_verified,omitempty"`
}

// Contact represents user's contact details.
type Contact struct {
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Website  string `bson:"website,omitempty" json:"website,omitempty"`
	Location string `bson:"location,omitempty" json:"location,omitempty"`
}

// SocialMediaURL represents the value and settings for socials.
type SocialMediaURL struct {
	Value     string `bson:"value,omitempty" json:"value,omitempty"`
	IsVisible bool   `bson:"is_visible,omitempty" json:"is_visible,omitempty"`
}

// Socials represents user's social details.
type Socials struct {
	GitHubURL    SocialMediaURL `bson:"github_url,omitempty" json:"github_url,omitempty"`
	LinkedInURL  SocialMediaURL `bson:"linkedin_url,omitempty" json:"linkedin_url,omitempty"`
	TwitterURL   SocialMediaURL `bson:"twitter_url,omitempty" json:"twitter_url,omitempty"`
	InstagramURL SocialMediaURL `bson:"instagram_url,omitempty" json:"instagram_url,omitempty"`
}

// A NotificationType represents notification settings.
type NotificationType struct {
	EmailEnabled bool `bson:"enable_email,omitempty" json:"enable_email,omitempty"`
	InAppEnabled bool `bson:"enable_in_app,omitempty" json:"enable_in_app,omitempty"`
}

// A Notifications represents user's notification details.
type Notifications struct {
	ID                           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	AllNotifications             NotificationType   `bson:"all_notifications,omitempty" json:"all_notifications,omitempty"`
	ProgramNotifications         NotificationType   `bson:"program_notifications,omitempty" json:"program_notifications,omitempty"`
	TaskNotifications            NotificationType   `bson:"task_notifications,omitempty" json:"task_notifications,omitempty"`
	ApprovalRequestNotifications NotificationType   `bson:"approval_request_notifications,omitempty" json:"approval_request_notifications,omitempty"`
	ReportsNotifications         NotificationType   `bson:"reports_notifications,omitempty" json:"reports_notifications,omitempty"`
	PostNotifications            NotificationType   `bson:"post_notifications,omitempty" json:"post_notifications,omitempty"`
	CommentsNotifications        NotificationType   `bson:"comments_notifications,omitempty" json:"comments_notifications,omitempty"`
	MentionsNotifications        NotificationType   `bson:"mentions_notifications,omitempty" json:"mentions_notifications,omitempty"`
	DirectMessageNotifications   NotificationType   `bson:"direct_message_notifications,omitempty" json:"direct_message_notifications,omitempty"`
	CommentsOnPostNotifications  NotificationType   `bson:"comments_on_post_notifications,omitempty" json:"comments_on_post_notifications,omitempty"`
}

// A Support represents support information.
type Support struct {
	Name  string `bson:"name,omitempty" json:"name,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
	Title string `bson:"title,omitempty" json:"title,omitempty"`
	Body  string `bson:"body,omitempty" json:"body,omitempty"`
}

// UserAction represents user action details
type UserAction struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email,omitempty"`
	SecretCode string             `bson:"secret_code,omitempty" json:"secret_code,omitempty"`
	ActionType string             `bson:"action_type,omitempty" json:"action_type,omitempty"`
	IsUsed     bool               `bson:"is_used,omitempty" json:"is_used,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	ExpiredAt  time.Time          `bson:"expired_at,omitempty" json:"expired_at,omitempty"`
}

// Faq represents faq details
type Faq struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Question string             `bson:"question,omitempty" json:"question,omitempty"`
	Answer   string             `bson:"answer,omitempty" json:"answer,omitempty"`
	Category string             `bson:"category,omitempty" json:"category,omitempty"`
}

// Discussion represents the data model for a discussion forum
type Discussion struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string             `bson:"title,omitempty" json:"title,omitempty"`
	Content   string             `bson:"content,omitempty" json:"content,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	OwnerID   primitive.ObjectID `bson:"owner_id,omitempty" json:"owner_id,omitempty"`
	Comments  []Comment          `bson:"comments,omitempty" json:"comments,omitempty"`
}

// Comment represents the data model for a comment
type Comment struct {
	OwnerID   primitive.ObjectID `bson:"owner_id,omitempty" json:"owner_id,omitempty"`
	FullName  string             `bson:"full_name,omitempty" json:"full_name,omitempty"`
	Content   string             `bson:"content,omitempty" json:"content,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
}
