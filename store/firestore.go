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

var storeClient *firestore.Client = nil

// InitializeFirestore initializes a Firestore client.
func InitializeFirestore() (*firestore.Client, error) {
	if storeClient != nil {
		return storeClient, nil
	}

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	projectID := os.Getenv("PROJECT_ID")

	ctx := context.Background()
	storeClient, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore client: %w", err)
	}

	return storeClient, nil
}
