package certificates

import (
	"github.com/openwurl/wurlwind/striketracker"
	"github.com/openwurl/wurlwind/striketracker/endpoints"
	"github.com/openwurl/wurlwind/striketracker/models"
)

/*
GET /api/v1/accounts/{account_hash}/certificates - List all certs on an account
POST /api/v1/accounts/{account_hash}/certificates - Upload a new certificate
DELETE /api/v1/accounts/{account_hash}/certificates/{certificate_id} - Delete a cert
GET/api/v1/accounts/{account_hash}/certificates/{certificate_id} - Get a certificate
PUT/api/v1/accounts/{account_hash}/certificates/{certificate_id} - Update a certificate (useful for expired certs)
GET/api/v1/accounts/{account_hash}/certificates/{certificate_id}/hosts - Get hosts for cert
*/

const path = "/certificates"

// Service describes the interaction with the origins API
type Service struct {
	client   *striketracker.Client
	Endpoint *endpoints.Endpoint
}

// New returns a new Origin Service
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
// Receives CertificateResponse
func (s *Service) List(accountHash string) (*models.CertificateResponse, error) {
	return nil, nil
}

// Get a certificate
//
// GET /api/v1/accounts/{account_hash}/certificates/{certificate_id}
//
// Receives Certificate
func (s *Service) Get(accountHash string, certificateID string) (*models.Certificate, error) {
	return nil, nil
}

// Hosts gets hosts for a certificate
//
// GET /api/v1/accounts/{account_hash}/certificates/{certificate_id}/hosts
//
// Receives CertificateHosts
func (s *Service) Hosts(accountHash string, certificateID string) (*models.CertificateHosts, error) {
	return nil, nil
}

// Upload a new certificate
//
// POST /api/v1/accounts/{account_hash}/certificates
//
// Sends Certificate
// Receives Certificate
func (s *Service) Upload(accountHash string, certificate *models.Certificate) (*models.Certificate, error) {
	return nil, nil
}

// Delete a certificate
//
// DELETE /api/v1/accounts/{account_hash}/certificates/{certificate_id}
//
// Pass in a models.Certificate with the ID set
func (s *Service) Delete(accountHash string, certificate *models.Certificate) error {
	return nil
}

// Update an existing certificate
//
// PUT /api/v1/accounts/{account_hash}/certificates/{certificate_id}
//
// Sends Certificate
// Receives Certificate
func (s *Service) Update(accountHash string, certificate *models.Certificate) (*models.Certificate, error) {
	return nil, nil
}
