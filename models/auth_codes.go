package models

import (
	"time"
)

type AuthCode struct {
	Code      string    `firestore:"code"`
	ClientID  string    `firestore:"client_id"`
	UserID    string    `firestore:"user_id"`
	ExpiresAt time.Time `firestore:"expires_at"`
	CreatedAt time.Time `firestore:"created_at"`
}
