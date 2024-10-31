package store

import (
	"context"
	"fmt"
	"oauth2-server/constants"
	"oauth2-server/models"

	"cloud.google.com/go/firestore"
)

// CreateToken creates a new token in Firestore.
func CreateToken(ctx context.Context, client *firestore.Client, newToken *models.Token) error {
	_, err := client.Collection(constants.FirestoreCollections.Tokens).Doc(newToken.AccessToken).Set(ctx, newToken)
	return err
}

// GetToken retrieves a token from Firestore by AccessToken.
func GetToken(ctx context.Context, client *firestore.Client, accessToken string) (*models.Token, error) {
	doc, err := client.Collection(constants.FirestoreCollections.Tokens).Doc(accessToken).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	var tokenData models.Token
	if err := doc.DataTo(&tokenData); err != nil {
		return nil, fmt.Errorf("failed to convert token data: %w", err)
	}

	return &tokenData, nil
}
