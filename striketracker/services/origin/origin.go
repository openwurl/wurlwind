// Package origin describes interactions with the Striketracker Origin service
//  c, err := striketracker.NewClientWithOptions(
//  	striketracker.WithApplicationID("DescriptiveApplicationName"),
//  	striketracker.WithDebug(true),
//  	striketracker.WithAuthorizationHeaderToken(authToken),
//  )
//  originService := origin.New(c)
//
// Context for early cancellation can be configured and passed in
//
//  ctx := context.Background()
//  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
//  defer cancel()
//
//  list, err := originService.List(ctx, accountHash)
//
// POST /api/v1/accounts/{account_hash}/origins - create new origin
// GET /api/v1/accounts/{account_hash}/origins - list all origins
// DELETE /api/v1/accounts/{account_hash}/origins/{origin_id} - delete
// GET /api/v1/accounts/{account_hash}/origins/{origin_id} - get one origin
// PUT /api/v1/accounts/{account_hash}/origins/{origin_id} - update origin
package origin

import (
	"context"
	"fmt"

	"github.com/openwurl/wurlwind/striketracker"
	"github.com/openwurl/wurlwind/striketracker/endpoints"
	"github.com/openwurl/wurlwind/striketracker/models"
	"github.com/openwurl/wurlwind/striketracker/services"
)

const path = "/origins"

// Service describes the interaction with the origins API
//  /api/v1/accounts/{account_hash}/origins
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
func (s *Service) Create(ctx context.Context, accountHash string, origin *models.Origin) (*models.Origin, error) {

	if err := origin.Validate(); err != nil {
		return nil, err
	}

	req, err := s.client.NewRequestContext(ctx, striketracker.POST, s.Endpoint.Format(accountHash), origin)
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
func (s *Service) Get(ctx context.Context, accountHash string, originID int) (*models.Origin, error) {

	endpoint := fmt.Sprintf("%s/%d", s.Endpoint.Format(accountHash), originID)

	req, err := s.client.NewRequestContext(ctx, striketracker.GET, endpoint, nil)
	if err != nil {
		return nil, err
	}

	origin := &models.Origin{}

	resp, err := s.client.DoRequest(req, origin)
	if err != nil {
		return nil, err
	}

	if err = services.ValidateResponse(resp); err != nil {

		// Catch any embedded errors in the body and add them to our response
		// This is difficult to make more generic and needs copied since
		// Response is inherited and not first class in the struct
		if respErr := origin.Err(err); respErr != nil {
			err = respErr
		}

		return nil, err
	}

	return origin, nil
}

// Delete an origin
//
// DELETE /api/v1/accounts/{account_hash}/origins/{origin_id}
//
func (s *Service) Delete(ctx context.Context, accountHash string, originID int) error {

	// Construct endpoint with originID
	endpoint := fmt.Sprintf("%s/%d", s.Endpoint.Format(accountHash), originID)

	req, err := s.client.NewRequestContext(ctx, striketracker.DELETE, endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := s.client.DoRequest(req, nil)
	if err != nil {
		return err
	}

	if err = services.ValidateResponse(resp); err != nil {
		fmt.Println("validating")
		return err
	}

	return nil
}

// Update an origin
//
// PUT /api/v1/accounts/{account_hash}/origins/{origin_id}
//
// Sends Origin
// Receives Origin
func (s *Service) Update(ctx context.Context, accountHash string, origin *models.Origin) (*models.Origin, error) {
	// Validate incoming origin payload
	if err := origin.Validate(); err != nil {
		return nil, err
	}

	// Construct endpoint with originID
	endpoint := fmt.Sprintf("%s/%d", s.Endpoint.Format(accountHash), origin.ID)

	req, err := s.client.NewRequestContext(ctx, striketracker.PUT, endpoint, origin)
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

// List returns all origins in the given account
//
// GET /api/v1/accounts/{account_hash}/origins - list all origins
//
// Receives OriginList
func (s *Service) List(ctx context.Context, accountHash string) (*models.OriginList, error) {

	ol := &models.OriginList{}

	req, err := s.client.NewRequestContext(ctx, striketracker.GET, s.Endpoint.Format(accountHash), nil)
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
