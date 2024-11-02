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

// Setters
func (c *Client) SetID(id string) {
	c.ID = id
}

func (c *Client) SetClientSecret(clientSecret string) {
	c.ClientSecret = clientSecret
}

func (c *Client) SetName(name string) {
	c.Name = name
}

func (c *Client) SetRedirectURIs(redirectURIs []string) {
	c.RedirectURIs = redirectURIs
}

func (c *Client) SetCreatedAt(createdAt time.Time) {
	c.CreatedAt = createdAt
}

func (c *Client) SetUpdatedAt(updatedAt time.Time) {
	c.UpdatedAt = updatedAt
}

// Getters
func (c *Client) GetID() string {
	return c.ID
}

func (c *Client) GetClientSecret() string {
	return c.ClientSecret
}

func (c *Client) GetName() string {
	return c.Name
}

func (c *Client) GetRedirectURIs() []string {
	return c.RedirectURIs
}

func (c *Client) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *Client) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}

// NewClient creates a new instance of Client
func NewClient() *Client {
	return &Client{}
}
