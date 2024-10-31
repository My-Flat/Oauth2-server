package models

import (
	"time"
)

type Token struct {
	AccessToken  string    `firestore:"access_token"`
	RefreshToken string    `firestore:"refresh_token"`
	ClientID     string    `firestore:"client_id"`
	UserID       string    `firestore:"user_id"`
	Scope        []string  `firestore:"scope"`
	ExpiresAt    time.Time `firestore:"expires_at"`
	Revoked      bool      `firestore:"revoked"`
	CreatedAt    time.Time `firestore:"created_at"`
}
