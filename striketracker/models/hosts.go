package models

import (
	"fmt"

	"github.com/openwurl/wurlwind/pkg/validation"
	validator "gopkg.in/go-playground/validator.v9"
)

const (
	defaultPath = "/"
	scopeCDS    = "CDS"
)

// Host defines the top level overview of a delivery host
type Host struct {
	Response
	Name        string             `json:"name" validate:"required"`
	HashCode    string             `json:"hashCode"`
	Type        string             `json:"type"`
	CreatedDate string             `json:"createdDate"`
	UpdatedDate string             `json:"updatedDate"`
	Services    []*DeliveryService `json:"services" validate:"required"`
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

// GetCDSScope returns the CDS "/" scope
func (h *Host) GetCDSScope() *Scope {
	var rootScope *Scope
	for _, scope := range h.Scopes {
		if scope.Platform == scopeCDS && scope.Path == defaultPath {
			rootScope = scope
		}
	}
	return rootScope
}

// NewDefaultHost returns a named host with CDS enabled
func NewDefaultHost(name string) *Host {
	h := &Host{
		Name: name,
	}
	h.Services = append(h.Services, ServiceCDS)
	return h
}

// HostList is a list of hosts
type HostList struct {
	Response
	List []*Host `json:"list"`
}

// Scope defines a delivery scope
type Scope struct {
	Response
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Platform    string `json:"platform" validate:"required"`
	Path        string `json:"path" validate:"required"`
	CreatedDate string `json:"createdDate,omitempty"`
	UpdatedDate string `json:"updatedDate,omitempty"`
}

// SetName sets the scope name of the configuration
func (s *Scope) SetName(name string) {
	s.Name = name
}

// GetName gets the scope name of the configuration
func (s *Scope) GetName() string {
	return s.Name
}

// GetIDString returns the scope ID as a string
func (s *Scope) GetIDString() string {
	return fmt.Sprintf("%d", s.ID)
}

// SetPlatform sets the platform type of the configuration
func (s *Scope) SetPlatform(platform string) {
	s.Platform = platform
}

// SetPath sets the scope path of the configuration
func (s *Scope) SetPath(path string) {
	s.Path = path
}

// Validate validates the struct data
func (s *Scope) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(s); err != nil {
		return err
	}
	return nil
}

// CloneHost is a type used when cloning an existing host
type CloneHost struct {
	Name      string   `json:"name" validate:"required"`
	Hostnames []string `json:"hostnames"`
}

// Validate validates the struct data
func (h *CloneHost) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(h); err != nil {
		return err
	}
	return nil
}
