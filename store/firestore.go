package store

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

const (
	ClientsCollection   = "clients"
	TokensCollection    = "tokens"
	AuthCodesCollection = "authcodes"
)

// InitializeFirestore initializes a Firestore client.
func InitializeFirestore() (*firestore.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	projectID := os.Getenv("PROJECT_ID")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore client: %w", err)
	}

	return client, nil
}
