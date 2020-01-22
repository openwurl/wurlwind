package models

// DeliveryService defines the type of billable service used by a host
// These are predefined by the CDN
type DeliveryService struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

var (
	// ServiceCDS is the HTTP caching service required on most site creations
	ServiceCDS = &DeliveryService{
		ID: 40,
	}
	// ServiceOriginShield is the reduced origin load shielding service
	ServiceOriginShield = &DeliveryService{
		ID: 62,
	}
)
