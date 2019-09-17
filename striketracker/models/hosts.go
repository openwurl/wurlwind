package models

// Host defines the top level overview of a delivery host
type Host struct {
	Name        string             `json:"name"`
	HashCode    string             `json:"hashCode"`
	Type        string             `json:"type"`
	CreatedDate string             `json:"createdDate"`
	UpdatedDate string             `json:"updatedDate"`
	Services    []*DeliveryService `json:"services"`
	Scopes      []*Scope           `json:"scopes"`
}

// Scope defines a delivery scope
type Scope struct {
	ID          int    `json:"id"`
	Platform    string `json:"platform"`
	Path        string `json:"path"`
	CreatedDate string `json:"createdDate"`
	UpdatedDate string `json:"updatedDate"`
}
