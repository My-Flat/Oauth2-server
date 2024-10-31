package models

import (
	"time"
)

type Client struct {
	ID           string    `firestore:"client_id"`
	ClientSecret string    `firestore:"client_secret"`
	Name         string    `firestore:"name"`
	RedirectURIs []string  `firestore:"redirect_uris"`
	CreatedAt    time.Time `firestore:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at"`
}
