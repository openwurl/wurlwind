// Package certificates surfaces actions for Uploading, deleting, and updating etc
// TLS certificates on striketracker
//  c, err := striketracker.NewClientWithOptions(
//  	striketracker.WithApplicationID("DescriptiveApplicationName"),
//  	striketracker.WithDebug(true),
//  	striketracker.WithAuthorizationHeaderToken(authToken),
//  )
//  certService := certificates.New(c)
//
// Context for early cancellation can be configured and passed in
//
//  ctx := context.Background()
//  ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
//  defer cancel()
//
//  list, err := certService.List(ctx, accountHash)
//
package certificates

import (
	"context"
	"fmt"

	"github.com/openwurl/wurlwind/striketracker"
	"github.com/openwurl/wurlwind/striketracker/endpoints"
	"github.com/openwurl/wurlwind/striketracker/models"
	"github.com/openwurl/wurlwind/striketracker/services"
)

const path = "/certificates"

// Service describes the interaction with the certificates API
type Service struct {
	client   *striketracker.Client
	Endpoint *endpoints.Endpoint
}

// New returns a new Certificates Service
func New(c *striketracker.Client) *Service {
	e := &endpoints.Endpoint{
		BasePath: endpoints.Accounts,
		Path:     path,
	}

	return &Service{
		Endpoint: e,
		client:   c,
	}

}

// List all certificates
//
// GET /api/v1/accounts/{account_hash}/certificates
//
// Returns models.CertificateResponse
func (s *Service) List(ctx context.Context, accountHash string) (*models.CertificateResponse, error) {
	cl := &models.CertificateResponse{}

	req, err := s.client.NewRequestContext(ctx, striketracker.GET, s.Endpoint.Format(accountHash), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(req, cl)
	if err != nil {
		return nil, err
	}

	if err = services.ValidateResponse(resp); err != nil {
		return nil, err
	}

	return cl, nil
}

// Get a certificate
//
// GET /api/v1/accounts/{account_hash}/certificates/{certificate_id}
//
// Accepts Certificate ID
//
// Returns models.Certificate
func (s *Service) Get(ctx context.Context, accountHash string, certificateID int) (*models.Certificate, error) {
	endpoint := fmt.Sprintf("%s/%d", s.Endpoint.Format(accountHash), certificateID)

	req, err := s.client.NewRequestContext(ctx, striketracker.GET, endpoint, nil)
	if err != nil {
		return nil, err
	}

	certificate := &models.Certificate{}

	resp, err := s.client.DoRequest(req, certificate)
	if err != nil {
		return nil, err
	}

	if err = services.ValidateResponse(resp); err != nil {

		// Catch any embedded errors in the body that supercede
		// validation errors and add them to our response
		// This is difficult to make more generic and needs copied since
		// Response is inherited and not first class in the struct
		if respErr := certificate.Error(); respErr != nil {
			err = respErr
		}

		return nil, err
	}

	return certificate, nil
}

// Hosts gets hosts for a certificate
//
// GET /api/v1/accounts/{account_hash}/certificates/{certificate_id}/hosts
//
// Receives models.CertificateHosts
//
// This is a weird one without the usually structured response
// so there is no error extraction
// It may need expanded once I come across a case with more than one Common Name
func (s *Service) Hosts(ctx context.Context, accountHash string, certificateID int) (*models.CertificateHosts, error) {
	endpoint := fmt.Sprintf("%s/%d/hosts", s.Endpoint.Format(accountHash), certificateID)
	req, err := s.client.NewRequestContext(ctx, striketracker.GET, endpoint, nil)
	if err != nil {
		return nil, err
	}

	certHostsResponse := &models.CertificateHostsResponse{}

	_, err = s.client.DoRequest(req, certHostsResponse)
	if err != nil {
		return nil, err
	}

	certHosts, err := certHostsResponse.Process()
	if err != nil {
		return nil, err
	}

	return certHosts, nil
}

// Upload a new certificate
//
// POST /api/v1/accounts/{account_hash}/certificates
//
// Accepts models.Certificate
//
// Returns models.Certificate
func (s *Service) Upload(ctx context.Context, accountHash string, certificate *models.Certificate) (*models.Certificate, error) {
	if err := certificate.Validate(); err != nil {
		return nil, err
	}

	req, err := s.client.NewRequestContext(ctx, striketracker.POST, s.Endpoint.Format(accountHash), certificate)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(req, certificate)
	if err != nil {
		return nil, err
	}

	if err = services.ValidateResponse(resp); err != nil {

		// Catch any embedded errors in the body and add them to our response
		if respErr := certificate.Error(); respErr != nil {
			err = respErr
		}

		return nil, err
	}

	return certificate, nil
}

// Delete a certificate
//
// DELETE /api/v1/accounts/{account_hash}/certificates/{certificate_id}
//
// Accepts Certificate ID
func (s *Service) Delete(ctx context.Context, accountHash string, certificateID int) error {
	// construct endpoint with certificateID
	endpoint := fmt.Sprintf("%s/%d", s.Endpoint.Format(accountHash), certificateID)

	req, err := s.client.NewRequestContext(ctx, striketracker.DELETE, endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := s.client.DoRequest(req, nil)
	if err != nil {
		return err
	}

	if err = services.ValidateResponse(resp); err != nil {
		return err
	}

	return nil
}

// Update an existing certificate
//
// PUT /api/v1/accounts/{account_hash}/certificates/{certificate_id}
//
// Accepts models.Certificate
//
// Returns updated models.Certificate
func (s *Service) Update(ctx context.Context, accountHash string, certificate *models.Certificate) (*models.Certificate, error) {
	// Validate incoming payload
	if err := certificate.Validate(); err != nil {
		return nil, err
	}

	// Construct
	endpoint := fmt.Sprintf("%s/%d", s.Endpoint.Format(accountHash), certificate.ID)

	// Build
	req, err := s.client.NewRequestContext(ctx, striketracker.PUT, endpoint, certificate)
	if err != nil {
		return nil, err
	}

	// Execute
	resp, err := s.client.DoRequest(req, certificate)
	if err != nil {
		return nil, err
	}

	// Validate response
	if err = services.ValidateResponse(resp); err != nil {

		// Catch any embedded errors in the body and add them to our response
		if respErr := certificate.Error(); respErr != nil {
			err = respErr
		}

		return nil, err
	}

	return certificate, nil
}
