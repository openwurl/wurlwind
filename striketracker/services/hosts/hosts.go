// Package hosts describes the interactions with the striketracker Hosts service
package hosts

import (
	"context"

	"github.com/openwurl/wurlwind/striketracker"
	"github.com/openwurl/wurlwind/striketracker/endpoints"
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
func (s *Service) Create(ctx context.Context) {

}

// Clone an existing host
func (s *Service) Clone(ctx context.Context) {

}

// Update a host
func (s *Service) Update(ctx context.Context) {

}

// Delete a host
func (s *Service) Delete(ctx context.Context) {

}

// Get a host
func (s *Service) Get(ctx context.Context) {

}

// List Hosts
func (s *Service) List(ctx context.Context) {

}
