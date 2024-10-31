package store

import (
	"context"
	"fmt"
	"oauth2-server/constants"
	"oauth2-server/models"

	"cloud.google.com/go/firestore"
)

// --- AuthCode functions ---

// CreateAuthCode creates a new authorization code in Firestore.
func CreateAuthCode(ctx context.Context, client *firestore.Client, newCode *models.AuthCode) error {
	_, err := client.Collection(constants.FirestoreCollections.AuthCodes).Doc(newCode.Code).Set(ctx, newCode)
	return err
}

// GetAuthCode retrieves an authorization code from Firestore by Code.
func GetAuthCode(ctx context.Context, client *firestore.Client, code string) (*models.AuthCode, error) {
	doc, err := client.Collection(constants.FirestoreCollections.AuthCodes).Doc(code).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth code: %w", err)
	}

	var codeData models.AuthCode
	if err := doc.DataTo(&codeData); err != nil {
		return nil, fmt.Errorf("failed to convert auth code data: %w", err)
	}

	return &codeData, nil
}
