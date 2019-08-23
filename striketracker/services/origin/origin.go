package origin

import (
	"fmt"

	"github.com/openwurl/wurlwind/striketracker"
	"github.com/openwurl/wurlwind/striketracker/endpoints"
	"github.com/openwurl/wurlwind/striketracker/models"
)

/*
POST /api/v1/accounts/{account_hash}/origins - create new origin
GET /api/v1/accounts/{account_hash}/origins - list all origins
DELETE /api/v1/accounts/{account_hash}/origins/{origin_id} - delete
GET /api/v1/accounts/{account_hash}/origins/{origin_id} - get one origin
PUT /api/v1/accounts/{account_hash}/origins/{origin_id} - update origin
*/

// Base /api/v1/accounts/{account_hash}/origins

const path = "/origins"

// Service describes the interaction with the origins API
type Service struct {
	client   *striketracker.Client
	Endpoint *endpoints.Endpoint
}

// New returns a new Origin Service
func New(c *striketracker.Client) *Service {
	e := &endpoints.Endpoint{
		BasePath: endpoints.Origins,
		Path:     path,
	}

	return &Service{
		Endpoint: e,
		client:   c,
	}

}

// List returns all origins in the given account
//
// GET /api/v1/accounts/{account_hash}/origins - list all origins
//
func (s *Service) List(accountHash string) (*models.OriginList, error) {

	ol := &models.OriginList{}

	req, err := s.client.CreateRequest(striketracker.GET, s.Endpoint.Format(accountHash), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(req, ol)
	if err != nil {
		return nil, err
	}

	// should validate the resp here, but just for now
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error: %d", resp.StatusCode)
	}

	return ol, nil
}
