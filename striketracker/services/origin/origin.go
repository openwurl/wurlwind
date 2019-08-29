package origin

import (
	"github.com/openwurl/wurlwind/striketracker"
	"github.com/openwurl/wurlwind/striketracker/endpoints"
	"github.com/openwurl/wurlwind/striketracker/models"
	"github.com/openwurl/wurlwind/striketracker/services"
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

// Create a new origin
//
// POST /api/v1/accounts/{account_hash}/origins
//
// Sends Origin
// Receives Origin
func (s *Service) Create(accountHash string, origin *models.Origin) (*models.Origin, error) {

	if err := origin.Validate(); err != nil {
		return nil, err
	}

	req, err := s.client.CreateRequest(striketracker.POST, s.Endpoint.Format(accountHash), origin)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(req, origin)
	if err != nil {
		return nil, err
	}

	if err = services.ValidateResponse(resp); err != nil {

		// Catch any embedded errors in the body and add them to our response
		if respErr := origin.Err(err); respErr != nil {
			err = respErr
		}

		return nil, err
	}

	return origin, nil
}

// Get an Origin
//
// GET /api/v1/accounts/{account_hash}/origins/{origin_id}
//
// Receives Origin
func (s *Service) Get(accountHash string, originID string) (*models.Origin, error) {
	return nil, nil
}

// Delete an origin
//
// DELETE /api/v1/accounts/{account_hash}/origins/{origin_id}
//
func (s *Service) Delete(accountHash string, originID string) error {
	return nil
}

// Update an origin
//
// PUT /api/v1/accounts/{account_hash}/origins/{origin_id}
//
// Sends Origin
// Receives Origin
func (s *Service) Update(accountHash string, origin *models.Origin) (*models.Origin, error) {
	return nil, nil
}

// List returns all origins in the given account
//
// GET /api/v1/accounts/{account_hash}/origins - list all origins
//
// Receives OriginList
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

	if err = services.ValidateResponse(resp); err != nil {
		return nil, err
	}

	return ol, nil
}
