package services

import (
	"context"
	"fmt"
	"time"

	"oauth2-server/models"
	"oauth2-server/store"

	"cloud.google.com/go/firestore"
)

// AuthService struct to hold dependencies like Firestore client
type AuthService struct {
	db *firestore.Client
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(db *firestore.Client) *AuthService {
	return &AuthService{db: db}
}

// RegisterClient registers a new client with the OAuth2 server
func (s *AuthService) RegisterClient(ctx context.Context, name string, redirectURIs []string) (string, string, error) {
	clientID := generateUniqueID()
	clientSecret := generateUniqueSecret()

	newClient := models.Client{
		ID:           clientID,
		ClientSecret: clientSecret,
		Name:         name,
		RedirectURIs: redirectURIs,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err := s.db.Collection(store.ClientsCollection).Doc(clientID).Set(ctx, newClient)
	if err != nil {
		return "", "", fmt.Errorf("failed to register client: %w", err)
	}

	return clientID, clientSecret, nil
}

// GenerateAuthCode generates an authorization code for a given client and user
func (s *AuthService) GenerateAuthCode(ctx context.Context, clientID, userID string, scopes []string, redirectURI string) (string, error) {
	authCode := generateUniqueCode()

	newAuthCode := models.AuthCode{
		Code:      authCode,
		ClientID:  clientID,
		UserID:    userID,
		ExpiresAt: time.Now().Add(10 * time.Minute),
		CreatedAt: time.Now(),
	}

	_, err := s.db.Collection(store.AuthCodesCollection).Doc(authCode).Set(ctx, newAuthCode)
	if err != nil {
		return "", fmt.Errorf("failed to generate auth code: %w", err)
	}

	return authCode, nil
}

// ExchangeAuthCodeForToken exchanges an auth code for an access and refresh token
func (s *AuthService) ExchangeAuthCodeForToken(ctx context.Context, authCode, clientID, redirectURI string) (string, string, error) {
	doc, err := s.db.Collection(store.AuthCodesCollection).Doc(authCode).Get(ctx)
	if err != nil {
		return "", "", fmt.Errorf("auth code not found: %w", err)
	}

	var code models.AuthCode
	doc.DataTo(&code)
	if code.ExpiresAt.Before(time.Now()) {
		return "", "", fmt.Errorf("auth code expired")
	}

	accessToken := generateUniqueToken()
	refreshToken := generateUniqueToken()

	newToken := models.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ClientID:     clientID,
		UserID:       code.UserID,
		Scope:        code.Scope,
		ExpiresAt:    time.Now().Add(1 * time.Hour),
		Revoked:      false,
		CreatedAt:    time.Now(),
	}

	_, err = s.db.Collection(store.TokensCollection).Doc(accessToken).Set(ctx, newToken)
	if err != nil {
		return "", "", fmt.Errorf("failed to save tokens: %w", err)
	}

	return accessToken, refreshToken, nil
}

// RevokeToken revokes an access or refresh token
func (s *AuthService) RevokeToken(ctx context.Context, token string) error {
	_, err := s.db.Collection(store.TokensCollection).Doc(token).Update(ctx, []firestore.Update{
		{Path: "revoked", Value: true},
	})
	if err != nil {
		return fmt.Errorf("failed to revoke token: %w", err)
	}
	return nil
}

// ValidateToken checks if an access token is valid and has the required scopes
func (s *AuthService) ValidateToken(ctx context.Context, accessToken string, requiredScopes []string) (bool, error) {
	doc, err := s.db.Collection(store.TokensCollection).Doc(accessToken).Get(ctx)
	if err != nil {
		return false, fmt.Errorf("token not found: %w", err)
	}

	var token models.Token
	doc.DataTo(&token)
	if token.ExpiresAt.Before(time.Now()) || token.Revoked {
		return false, fmt.Errorf("token expired or revoked")
	}

	if !containsAllScopes(token.Scope, requiredScopes) {
		return false, fmt.Errorf("insufficient token scope")
	}

	return true, nil
}

// RefreshAccessToken refreshes an access token using a refresh token
func (s *AuthService) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	doc, err := s.db.Collection(store.TokensCollection).Where("refresh_token", "==", refreshToken).Documents(ctx).Next()
	if err != nil {
		return "", fmt.Errorf("refresh token not found: %w", err)
	}

	var token models.Token
	doc.DataTo(&token)
	if token.ExpiresAt.Before(time.Now()) || token.Revoked {
		return "", fmt.Errorf("refresh token expired or revoked")
	}

	newAccessToken := generateUniqueToken()

	_, err = s.db.Collection(store.TokensCollection).Doc(newAccessToken).Set(ctx, models.Token{
		AccessToken:  newAccessToken,
		RefreshToken: refreshToken,
		ClientID:     token.ClientID,
		UserID:       token.UserID,
		Scope:        token.Scope,
		ExpiresAt:    time.Now().Add(1 * time.Hour),
		CreatedAt:    time.Now(),
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate new access token: %w", err)
	}

	return newAccessToken, nil
}
