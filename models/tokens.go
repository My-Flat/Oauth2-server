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

// Setters

func (t *Token) SetAccessToken(accessToken string) {
	t.AccessToken = accessToken
}

func (t *Token) SetRefreshToken(refreshToken string) {
	t.RefreshToken = refreshToken
}

func (t *Token) SetClientID(clientID string) {
	t.ClientID = clientID
}

func (t *Token) SetUserID(userID string) {
	t.UserID = userID
}

func (t *Token) SetScope(scope []string) {
	t.Scope = scope
}

func (t *Token) SetExpiresAt(expiresAt time.Time) {
	t.ExpiresAt = expiresAt
}

func (t *Token) SetRevoked(revoked bool) {
	t.Revoked = revoked
}

func (t *Token) SetCreatedAt(createdAt time.Time) {
	t.CreatedAt = createdAt
}

// Getters

func (t *Token) GetAccessToken() string {
	return t.AccessToken
}

func (t *Token) GetRefreshToken() string {
	return t.RefreshToken
}

func (t *Token) GetClientID() string {
	return t.ClientID
}

func (t *Token) GetUserID() string {
	return t.UserID
}

func (t *Token) GetScope() []string {
	return t.Scope
}

func (t *Token) GetExpiresAt() time.Time {
	return t.ExpiresAt
}

func (t *Token) GetRevoked() bool {
	return t.Revoked
}

func (t *Token) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func NewToken() *Token {
	return &Token{}
}
