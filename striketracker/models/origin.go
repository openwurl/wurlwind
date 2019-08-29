package models

/*
POST /api/v1/accounts/{account_hash}/origins - create new origin
GET /api/v1/accounts/{account_hash}/origins - list all origins
DELETE /api/v1/accounts/{account_hash}/origins/{origin_id} - delete
GET /api/v1/accounts/{account_hash}/origins/{origin_id} - get one origin
PUT /api/v1/accounts/{account_hash}/origins/{origin_id} - update origin
*/

import (
	"github.com/openwurl/wurlwind/pkg/validation"

	validator "gopkg.in/go-playground/validator.v9"
)

// OriginList unwraps a list of origins from the API
type OriginList struct {
	List []Origin `json:"list"`
}

// Origin is the central type for a highwind CDN origin
type Origin struct {
	Response
	// Required
	Name     string `json:"name" validate:"required"`
	Hostname string `json:"hostname" validate:"required,domain"`
	Port     int    `json:"port" validate:"required,oneof=80 443 8080 1935"` // supports 80, 443, 8080, 1935

	// Optional
	ID                           int    `json:"id,omitempty"`
	AuthenticationType           string `json:"authenticationType,omitempty"` // NONE or BASIC  validate:"oneof=NONE BASIC"
	CertificateCN                string `json:"certificateCN,omitempty"`      // CertificateCN Common name to validate if VerifyCertificate
	CreatedDate                  string `json:"createdDate,omitempty"`
	ErrorCacheTTLSeconds         int    `json:"errorCacheTTLSeconds,omitempty"`         // DNS Timeout
	MaxConnectionsPerEdge        int    `json:"maxConnectionsPerEdge,omitempty"`        // If enabled, the maximum number of concurrent connection any single edge will make to the origin
	MaxConnectionsPerEdgeEnabled int    `json:"maxConnectionsPerEdgeEnabled,omitempty"` // Indicates if the CDN should limit the number of connections each edge should make when pulling content
	MaximumOriginPullSeconds     int    `json:"maximumOriginPullSeconds,omitempty"`
	MaxRetryCount                int    `json:"maxRetryCount,omitempty"`
	OriginCacheHeaders           string `json:"originCacheHeaders,omitempty"` // Access-Control-Allow-Origin
	OriginDefaultKeepAlive       int    `json:"originDefaultKeepAlive,omitempty"`
	OriginPullHeaders            string `json:"originPullHeaders,omitempty"`
	OriginPullNegLinger          string `json:"originPullNegLinger,omitempty"`
	Path                         string `json:"path,omitempty" validate:"path"`
	RequestTimeoutSeconds        int    `json:"requestTimeoutSeconds,omitempty"` // Default 15
	SecurePort                   int    `json:"securePort,omitempty"`
	Type                         string `json:"type,omitempty"` // Default EXTERNAL
	UpdatedDate                  string `json:"updatedDate,omitempty"`
	VerifyCertificate            bool   `json:"verifyCertificate,omitempty"`
}

// Validate validates the struct data
func (o *Origin) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(o); err != nil {
		return err
	}
	return nil
}
