// Package configuration interacts with the cofiguration service of the Striketracker API
package configuration

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/imdario/mergo"
	"github.com/openwurl/wurlwind/striketracker"
	"github.com/openwurl/wurlwind/striketracker/endpoints"
	"github.com/openwurl/wurlwind/striketracker/models"
	"github.com/openwurl/wurlwind/striketracker/services"
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
// Accepts models.Configuration and hostHash
//
// Returns an updated models.Configuration
func (s *Service) Create(ctx context.Context, accountHash string, hostHash string, scope *models.Configuration) (*models.Configuration, error) {
	if err := scope.Validate(); err != nil {
		return nil, err
	}

	spew.Dump(scope)

	endpoint := fmt.Sprintf("%s/%s/configuration/scopes", s.Endpoint.Format(accountHash), hostHash)

	req, err := s.client.NewRequestContext(ctx, striketracker.POST, endpoint, scope)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(req, scope)
	if err != nil {
		return nil, err
	}

	if err = services.ValidateResponse(resp); err != nil {

		// Catch any embedded errors in the body and add them to our response
		if respErr := scope.Error(); respErr != nil {
			err = respErr
		}

		return nil, err
	}

	return scope, nil
}

// Update a scope's configuration
//
// PUT /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}
//
// Accepts models.Configuration
//
// Returns an updated models.Configuration
func (s *Service) Update(ctx context.Context, accountHash string, hostHash string, scopeID int, config *models.Configuration) (*models.Configuration, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Fetch upstream configuration then merge our changes to it
	// Changes should be explicit
	origin, err := s.Get(ctx, accountHash, hostHash, scopeID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get upstream configuration to merge: %v", err)
	}
	err = mergo.Merge(origin, config)
	if err != nil {
		return nil, fmt.Errorf("Failed to merge upstream and given configuration")
	}

	endpoint := fmt.Sprintf("%s/%s/configuration/%d", s.Endpoint.Format(accountHash), accountHash, scopeID)

	req, err := s.client.NewRequestContext(ctx, striketracker.PUT, endpoint, config)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(req, config)
	if err != nil {
		return nil, err
	}

	if err = services.ValidateResponse(resp); err != nil {

		// Catch any embedded errors in the body and add them to our response
		if respErr := config.Error(); respErr != nil {
			err = respErr
		}

		return nil, err
	}

	return config, nil
}

// Get a scope's configuration
//
// GET /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}
//
// Accepts scope id
//
// Returns *models.Configuration
func (s *Service) Get(ctx context.Context, accountHash string, hostHash string, scopeID int) (*models.Configuration, error) {
	endpoint := fmt.Sprintf("%s/%s/configuration/%d", s.Endpoint.Format(accountHash), hostHash, scopeID)

	req, err := s.client.NewRequestContext(ctx, striketracker.GET, endpoint, nil)
	if err != nil {
		return nil, err
	}

	scopeConfig := &models.Configuration{}

	resp, err := s.client.DoRequest(req, scopeConfig)
	if err != nil {
		return nil, err
	}

	if err = services.ValidateResponse(resp); err != nil {

		// Catch any embedded errors in the body and add them to our response
		// This is difficult to make more generic and needs copied since
		// Response is inherited and not first class in the struct
		if respErr := scopeConfig.Error(); respErr != nil {
			err = respErr
		}

		return nil, err
	}

	if &scopeConfig.Scope == nil {
		return nil, fmt.Errorf("Scope payload did not flesh")
	}

	return scopeConfig, nil
}

// Delete a configuration scope
//
// DELETE /api/v1/accounts/{account_hash}/hosts/{host_hash}/configuration/{scope_id}?force={forceDelete}
//
// Accepts ScopeID and forceDelete
//
// Returns error
func (s *Service) Delete(ctx context.Context, accountHash string, hostHash string, scopeID int, forceDelete bool) error {
	endpoint := fmt.Sprintf("%s/%s/configuration/%d", s.Endpoint.Format(accountHash), hostHash, scopeID)

	req, err := s.client.NewRequestContext(ctx, striketracker.GET, endpoint, nil)
	if err != nil {
		return err
	}

	req = striketracker.AddRequestParameter(req, "force", fmt.Sprintf("%t", forceDelete))

	resp, err := s.client.DoRequest(req, nil)
	if err != nil {
		return err
	}

	if err = services.ValidateResponse(resp); err != nil {
		return err
	}

	return nil
}
