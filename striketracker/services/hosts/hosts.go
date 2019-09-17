// Package hosts describes the interactions with the striketracker Hosts service
package hosts

import (
	"context"

	"github.com/openwurl/wurlwind/striketracker"
	"github.com/openwurl/wurlwind/striketracker/endpoints"
	"github.com/openwurl/wurlwind/striketracker/models"
	"github.com/openwurl/wurlwind/striketracker/services"
)

/*
POST /api/v1/accounts/{account_hash}/hosts Create a new delivery host
GET /api/v1/accounts/{account_hash}/hosts List delivery hosts for the specified account
POST /api/v1/accounts/{account_hash}/hosts/{host_hash} Clone an existing delivery host
DELETE /api/v1/accounts/{account_hash}/hosts/{host_hash} Delete a delivery host
GET /api/v1/accounts/{account_hash}/hosts/{host_hash} Get a delivery host
PUT /api/v1/accounts/{account_hash}/hosts/{host_hash} Update a delivery host
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
		BasePath: endpoints.Hosts,
		Path:     path,
	}
	return &Service{
		Endpoint: e,
		client:   c,
	}
}

// Create a new host
//
// POST /api/v1/accounts/{account_hash}/hosts
//
// Accepts models.Host with services defined
//
// Returns an updated models.Host
func (s *Service) Create(ctx context.Context, accountHash string, host *models.Host) (*models.Host, error) {

	if err := host.Validate(); err != nil {
		return nil, err
	}

	req, err := s.client.NewRequestContext(ctx, striketracker.POST, s.Endpoint.Format(accountHash), host)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(req, host)
	if err != nil {
		return nil, err
	}

	if err = services.ValidateResponse(resp); err != nil {
		// Catch any embedded errors in the body and add them to our response
		if respErr := host.Error(); respErr != nil {
			err = respErr
		}

		return nil, err
	}

	return host, nil
}

// Clone an existing host
func (s *Service) Clone(ctx context.Context, accountHash string, hostHash string, cloneHost *models.CloneHost) (*models.Host, error) {

	return nil, nil
}

// Update a host
func (s *Service) Update(ctx context.Context, accountHash string, hostHash string, host *models.Host) (*models.Host, error) {

	if err := host.Validate(); err != nil {
		return nil, err
	}

	return host, nil
}

// Delete a host
func (s *Service) Delete(ctx context.Context, accountHash string, hostHash string) error {
	return nil
}

// Get a host
func (s *Service) Get(ctx context.Context, accountHash string, hostHash string) (*models.Host, error) {
	return nil, nil
}

// List Hosts
func (s *Service) List(ctx context.Context, accountHash string, recursive bool) (*models.HostList, error) {
	return nil, nil
}
