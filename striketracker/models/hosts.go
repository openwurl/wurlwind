package models

import (
	"github.com/openwurl/wurlwind/pkg/validation"
	validator "gopkg.in/go-playground/validator.v9"
)

// Host defines the top level overview of a delivery host
type Host struct {
	Response
	Name        string             `json:"name"`
	HashCode    string             `json:"hashCode"`
	Type        string             `json:"type"`
	CreatedDate string             `json:"createdDate"`
	UpdatedDate string             `json:"updatedDate"`
	Services    []*DeliveryService `json:"services"`
	Scopes      []*Scope           `json:"scopes"`
}

// Validate validates the struct data
func (h *Host) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(h); err != nil {
		return err
	}
	return nil
}

// HostList is a list of hosts
type HostList struct {
	List []*Host `json:"list"`
}

// Scope defines a delivery scope
type Scope struct {
	ID          int    `json:"id"`
	Platform    string `json:"platform"`
	Path        string `json:"path"`
	CreatedDate string `json:"createdDate"`
	UpdatedDate string `json:"updatedDate"`
}

// CloneHost is a type used when cloning an existing host
type CloneHost struct {
	Name      string   `json:"name"`
	Hostnames []string `json:"hostnames"`
}
