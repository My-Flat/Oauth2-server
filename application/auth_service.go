package application

import (
	"context"
	"fmt"
	"time"

	"oauth2-server/store"

	"cloud.google.com/go/firestore"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"

	ext_store "github.com/go-oauth2/oauth2/v4/store"

	"github.com/google/uuid"
)

var a = ext_store.TokenStore{} // TODO: delete this line

type AuthService struct {
	oauthServer *server.Server
	clientStore *firestore.Client
}

func NewAuthService(clientStore *store.FirestoreClient) *AuthService {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore()) // You can use your Firestore client store here eventually

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), oauth2.SigningMethodHS512))

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	return &AuthService{
		oauthServer: srv,
		clientStore: clientStore,
	}
}

func (s *AuthService) Authorize(ctx context.Context, req *AuthorizeRequest) (*AuthorizeResponse, error) {
	// 1. Validate the authorization request (req)
	// 2. Authenticate the user (if needed) - This will depend on your authentication mechanism
	// 3. Check if the client is authorized

	// Assuming you have the authenticated user's ID
	userID := "user-id-123" // Replace with actual user ID

	// Generate an authorization code
	code, err := s.oauthServer.GetAuthorizeToken(ctx, oauth2.ResponseTypeCode, &models.Token{
		ClientID:    req.ClientID,
		UserID:      userID,
		RedirectURI: req.RedirectURI,
		Scope:       req.Scope,
		ExpiresIn:   time.Hour, // Customize as needed
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate authorization code: %w", err)
	}

	// 4. Build the authorization response
	return &AuthorizeResponse{
		Code:        code.GetCode(),
		RedirectURI: req.RedirectURI,
		State:       req.State,
	}, nil
}

func (s *AuthService) GetToken(ctx context.Context, req *TokenRequest) (*TokenResponse, error) {
	// 1. Validate the token request (req)

	// 2. Authenticate the client (if needed) - This might involve checking client credentials
	//    For simplicity, we're skipping client authentication here

	// 3. Exchange the authorization code for an access token
	token, err := s.oauthServer.GetAccessToken(ctx, oauth2.GrantTypeAuthorizationCode, req.ToMap())
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	// 4. Build the token response
	return &TokenResponse{
		AccessToken:  token.GetAccess(),
		RefreshToken: token.GetRefresh(),
		TokenType:    token.GetTokenType(),
		ExpiresIn:    token.GetAccessExpiresIn(),
	}, nil
}

// ... (other application logic functions) ...

// Define request/response structs
type AuthorizeRequest struct {
	ResponseType string
	ClientID     string
	RedirectURI  string
	Scope        string
	State        string
}

type AuthorizeResponse struct {
	Code        string
	RedirectURI string
	State       string
}

type TokenRequest struct {
	GrantType    string
	Code         string
	RedirectURI  string
	ClientID     string
	ClientSecret string
}

func (r *TokenRequest) ToMap() (m map[string]string) {
	m = make(map[string]string)
	m["grant_type"] = r.GrantType
	m["code"] = r.Code
	m["redirect_uri"] = r.RedirectURI
	m["client_id"] = r.ClientID
	m["client_secret"] = r.ClientSecret
	return
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`

	ExpiresIn time.Duration `json:"expires_in"`
}

// Helper function to generate a client ID
func generateClientID() string {
	return uuid.New().String()
}
