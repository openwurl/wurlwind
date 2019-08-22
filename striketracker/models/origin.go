package models

/*
POST /api/v1/accounts/{account_hash}/origins - create new origin
GET /api/v1/accounts/{account_hash}/origins - list all origins
DELETE /api/v1/accounts/{account_hash}/origins/{origin_id} - delete
GET /api/v1/accounts/{account_hash}/origins/{origin_id} - get one origin
PUT /api/v1/accounts/{account_hash}/origins/{origin_id} - update origin
*/

import (
	"github.com/wurlinc/hls-config/pkg/validate"
	validator "gopkg.in/go-playground/validator.v9"
)

// OriginList unwraps a list of origins from the API
type OriginList struct {
	List []Origin `json:"list"`
}

// Origin is the central type for a highwind CDN origin
type Origin struct {
	// Required
	Name     string `json:"name" validate:"required"`
	Hostname string `json:"hostname" validate:"required"`
	Port     int    `json:"port" validate:"required"`
	// Optional
	ID                           int    `json:"id,omitempty"`
	AuthenticationType           string `json:"authenticationType,omitempty"`
	CertificateCN                string `json:"certificateCN,omitempty"`
	CreatedDate                  string `json:"createdDate,omitempty"`
	ErrorCacheTTLSeconds         int    `json:"errorCacheTTLSeconds,omitempty"`
	MaxConnectionsPerEdge        int    `json:"maxConnectionsPerEdge,omitempty"`
	MaxConnectionsPerEdgeEnabled int    `json:"maxConnectionsPerEdgeEnabled,omitempty"`
	MaximumOriginPullSeconds     int    `json:"maximumOriginPullSeconds,omitempty"`
	MaxRetryCount                int    `json:"maxRetryCount,omitempty"`
	OriginCacheHeaders           string `json:"originCacheHeaders,omitempty"`
	OriginDefaultKeepAlive       int    `json:"originDefaultKeepAlive,omitempty"`
	OriginPullHeaders            string `json:"originPullHeaders,omitempty"`
	OriginPullNegLinger          string `json:"originPullNegLinger,omitempty"`
	Path                         string `json:"path,omitempty"`
	RequestTimeoutSeconds        int    `json:"requestTimeoutSeconds,omitempty"`
	SecurePort                   int    `json:"securePort,omitempty"`
	Type                         string `json:"type,omitempty"`
	UpdatedDate                  string `json:"updatedDate,omitempty"`
	VerifyCertificate            bool   `json:"verifyCertificate,omitempty"`
}

// Validate validates the struct data
func (o *Origin) Validate() error {
	v := validate.NewValidator(validator.New())
	if err := v.Validate(o); err != nil {
		return err
	}
	return nil
}
