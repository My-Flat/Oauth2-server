package store

import (
	// ... other imports ...
	"context"
	"fmt"
	"oauth2-server/constants"
	"oauth2-server/models"

	"cloud.google.com/go/firestore"
)

func CreateClient(ctx context.Context, client *firestore.Client, newClient *models.Client) error {
	_, err := client.Collection(constants.FirestoreCollections.Clients).Doc(newClient.ID).Set(ctx, newClient)
	return err
}

// GetClient retrieves a client from Firestore by ID.
func GetClient(ctx context.Context, client *firestore.Client, clientID string) (*models.Client, error) {
	doc, err := client.Collection(constants.FirestoreCollections.Clients).Doc(clientID).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %w", err)
	}

	var clientData models.Client
	if err := doc.DataTo(&clientData); err != nil {
		return nil, fmt.Errorf("failed to convert client data: %w", err)
	}

	return &clientData, nil
}
