package models

import (
	"github.com/openwurl/wurlwind/pkg/validation"
	validator "gopkg.in/go-playground/validator.v9"
)

/*

	Configuration for hosts

*/

// Configuration defines a high level scope configuration for a delivery hash
type Configuration struct {
	Response
	*Scope                     `json:"scope,omitempty"`
	Hostname                   []*ScopeHostname              `json:"hostname,omitempty"`
	OriginPullHost             *OriginPullHost               `json:"originPullHost,omitempty"`
	OriginPullCacheExtension   *OriginPullCacheExtension     `json:"originPullCacheExtension,omitempty"`
	OriginPullPolicy           []*OriginPullPolicy           `json:"originPullPolicy,omitempty"`
	OriginRequestModification  []*OriginRequestModification  `json:"originRequestModification,omitempty"`
	OriginResponseModification []*OriginResponseModification `json:"originResponseModification,omitempty"`
	ClientRequestModification  []*ClientRequestModification  `json:"clientRequestModification,omitempty"`
	ClientResponseModification []*ClientResponseModification `json:"clientResponseModification,omitempty"`
}

// NewConfiguration returns an empty configuration
func NewConfiguration() *Configuration {
	return &Configuration{}
}

// Validate validates the struct data
func (c *Configuration) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(c); err != nil {
		return err
	}
	return nil
}

/*

	Configuration for creating a new host

*/

// NewHostConfiguration defines a high level scope configuration for creating a new delivery hash
type NewHostConfiguration struct {
	Response
	Name     string `json:"name"`
	Platform string `json:"platform" validate:"required"`
	Path     string `json:"path" validate:"required"`
	ID       int    `json:"id"`
}

// Validate validates the struct data
func (c *NewHostConfiguration) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(c); err != nil {
		return err
	}
	return nil
}

// NewHostConfigurationFromState returns a baseline configuration from terraform state for create
func NewHostConfigurationFromState(state map[string]interface{}) (*NewHostConfiguration, error) {
	config := &NewHostConfiguration{}

	if state["platform"] != nil {
		config.Platform = state["platform"].(string)
	} else {
		config.Platform = "CDS"
	}

	if state["name"] != nil {
		config.Name = state["name"].(string)
	}

	if state["path"] != nil {
		config.Path = state["path"].(string)
	}

	return config, config.Validate()
}
