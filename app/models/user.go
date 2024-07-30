package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID  `json:"id" groups:"public"`
	Name       string     `json:"name" groups:"public"`
	Email      string     `json:"email" groups:"public"`
	Address    string     `json:"title" groups:"public"`
	PhoneNo    string     `json:"phone_no" groups:"public"`
	Password   string     `json:"password"`
	VerifiedAt *time.Time `json:"verified_at" groups:"public"`
	Channel    string     `json:"channel"`
	CreatedAt  time.Time  `json:"created_at" groups:"public"`
	UpdatedAt  time.Time  `json:"updated_at" `
}

const UserRegistrationKeyVerificationCode = "user-verfication-registration"
const AuthInternalApps = "internal"
const AuthGoogleApps = "google"
const AuthFacebookApps = "facebook"
