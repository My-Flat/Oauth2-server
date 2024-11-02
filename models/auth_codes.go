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

// Setters
func (a *AuthCode) SetCode(code string) {
	a.Code = code
}

func (a *AuthCode) SetClientID(clientID string) {
	a.ClientID = clientID
}

func (a *AuthCode) SetUserID(userID string) {
	a.UserID = userID
}

func (a *AuthCode) SetExpiresAt(expiresAt time.Time) {
	a.ExpiresAt = expiresAt
}

func (a *AuthCode) SetCreatedAt(createdAt time.Time) {
	a.CreatedAt = createdAt
}

// Getters

func (a *AuthCode) GetCode() string {
	return a.Code
}

func (a *AuthCode) GetClientID() string {
	return a.ClientID
}

func (a *AuthCode) GetUserID() string {
	return a.UserID
}

func (a *AuthCode) GetExpiresAt() time.Time {
	return a.ExpiresAt
}

func (a *AuthCode) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func NewAuthCode() *AuthCode {
	return &AuthCode{}
}
