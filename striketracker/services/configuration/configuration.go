// Package configuration interacts with the cofiguration service of the Striketracker API
package configuration

import (
	"context"

	"github.com/openwurl/wurlwind/striketracker"
	"github.com/openwurl/wurlwind/striketracker/endpoints"
	"github.com/openwurl/wurlwind/striketracker/models"
)

/*
GET /api/v1/accounts/{account_hash}/graph Get the configuration graph for the selected account
GET /api/v1/accounts/{account_hash}/hostnames List the hostnames that exist for an account
POST /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/scopes Create a new configuration scope for a given host
GET /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/scopes List the scopes at which configuration exists for a given host
DELETE /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id} Delete a configuration scope
GET /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id} Get host configuration at a certain scope
PUT /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id} Update host configuration at a certain scope
GET /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}/{configuration_receipt_id} Check on configuration update status
GET /api/v1/configuration List the configuration types that this API supports
*/

const path = "/hosts"

// Service describes the interaction with the hosts API
// and contains the instantiated client
type Service struct {
	client   *striketracker.Client
	Endpoint *endpoints.Endpoint
}

// New returns a new Hosts service
func New(c *striketracker.Client) *Service {
	e := &endpoints.Endpoint{
		BasePath: endpoints.Configuration,
		Path:     path,
	}
	return &Service{
		Endpoint: e,
		client:   c,
	}
}

// Create a new Scope
//
// POST /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/scopes
//
// Accepts models.Scope
//
// Returns an updated models.Scope
func (s *Service) Create(ctx context.Context, accountHash string, scope *models.Scope) (*models.Scope, error) {
	return nil, nil
}

// Update a scope's configuration
//
// PUT /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}
//
// Accepts models.Configuration
//
// Returns an updated models.Configuration
func (s *Service) Update(ctx context.Context, accountHash string, ScopeID int, config *models.Configuration) (*models.Configuration, error) {
	return nil, nil
}

// Get a scope's configuration
//
// GET /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}
//
// Accepts scope id
//
// Returns *models.Scope
func (s *Service) Get(ctx context.Context, accountHash string, scopeID int) (*models.Configuration, error) {
	return nil, nil
}

// Delete a configuration scope
//
// DELETE /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}
//
// Accepts ScopeID and forceDelete
//
// Returns error
func (s *Service) Delete(ctx context.Context, accountHash string, scopeID int, forceDelete bool) error {
	return nil
}
