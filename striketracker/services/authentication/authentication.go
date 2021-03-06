// Package authentication manages the creation and destruction
// of long-lived access tokens
package authentication

import (
	"context"

	"github.com/openwurl/wurlwind/striketracker"
	"github.com/openwurl/wurlwind/striketracker/endpoints"
	"github.com/openwurl/wurlwind/striketracker/models"
	"github.com/openwurl/wurlwind/striketracker/services"
)

/*
POST /api/v1/accounts/{account_hash}/users/{user_id}/tokens
GET /api/v1/accounts/{account_hash}/users/{user_id}/tokens
DELETE /api/v1/accounts/{account_hash}/users/{user_id}/tokens/{token_id}
NOT IMPLEMENTING POST /auth/token
*/

const path = "/tokens"

// Service describes the interaction with the auth API
type Service struct {
	client   *striketracker.Client
	Endpoint *AuthEndpoint
}

// New returns a new Auth Service
func New(c *striketracker.Client) *Service {
	e := &AuthEndpoint{
		&endpoints.Endpoint{
			BasePath: endpoints.ACCOUNTS,
			Path:     path,
		},
	}

	return &Service{
		Endpoint: e,
		client:   c,
	}
}

// Create an API token with infinite expiration
//
// POST /api/v1/accounts/{account_hash}/users/{user_id}/tokens
//
// Sends AccountHash, UserID, APITokenRequest
// Receives Authentication
func (s *Service) Create(ctx context.Context, accountHash string, userID string, password string, application string) (*models.Authentication, error) {

	payload := &models.CreateTokenRequest{
		AccountHash: accountHash,
		UserID:      userID,
		APITokenRequest: &models.APITokenRequest{
			Application: application,
			Password:    password,
		},
	}

	answer := &models.Authentication{}

	req, err := s.client.NewRequestContext(ctx, striketracker.POST, s.Endpoint.formatUser(accountHash, userID), payload)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(req, answer)
	if err != nil {
		return nil, err
	}

	if err = services.ValidateResponse(resp); err != nil {
		return nil, err
	}

	return answer, nil
}

// List Fetch all tokens associated with the user
//
// GET /api/v1/accounts/{account_hash}/users/{user_id}/tokens
//
// Sends AccountHash, UserID
// Receives AccessTokenList
func (s *Service) List(ctx context.Context, accountHash string, userID string) (*models.AccessTokenList, error) {
	return nil, nil
}

// Delete a token
//
// DELETE /api/v1/accounts/{account_hash}/users/{user_id}/tokens/{token_id}
//
// Sends AccountHash, UserID, TokenID
// Receives status
func (s *Service) Delete(ctx context.Context, accountHash string, userID string, token string) error {
	return nil
}
