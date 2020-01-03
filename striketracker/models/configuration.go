package models

import (
	"github.com/openwurl/wurlwind/pkg/validation"
	validator "gopkg.in/go-playground/validator.v9"
)

/*

	Configuration for hosts

*/

/*
	Missing:
		Delivery:
			Compression
			StaticHeader
			HTTPMethods
			GzipOriginPull

		Cache:
			CacheKeyModification
		Security:
			Geographic Restrictions
			IP Address Restrictions
			Referrer Restrictions
			Url Signing
		Reporting:
			OriginPullLogs
			AccessLogs
		Origin:
			OriginPullProtocol
			FileSegmentation
			OriginPersistentConnections
			OriginPull


	Missing Next Pass:
		Delivery:
			Bandwidth Rate Limiting
			Force Downloads
			Custom Mime Types
			Edge Responses
			Media Delivery
			Custom HTTP Response Headers
			Error Redirects
			Dynamic Files (Robots)
			TLS Configuration
		Cache:
			Bypass Cache
			Dynamic File Versioning
		Security:
			Basic Auth
			ASYMMETRIC URL signing
			IQIYI URL Signing
		Origin:
			Custom Client Identification
			Uncategorized (signed origin pull, fail safe pull, resume download etc)
		Reporting:
			Request Receipts
			Signed AWS Post

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
	Compression                *Compression                  `json:"compression,omitempty"`
	StaticHeader               []*StaticHeader               `json:"staticHeader,omitempty"`
	HTTPMethods                []*HTTPMethods                `json:"httpMethods,omitempty"`
	//CustomMimeType             []*CustomMimeType             `json:"customMimeType,omitempty"`
	//BandwidthLimit             *BandwidthLimit               `json:"bandWidthLimit,omitempty"`
	//BandwidthRateLimit         *BandwidthRateLimit           `json:"bandwidthRateLimit,omitempty"`
	//ContentDispositionByHeader []*ContentDispositionByHeader `json:"contentDispositionByHeader,omitempty"`
	//DynamicCacheRule           []*DynamicCacheRule           `json:"dynamicCacheRule,omitempty"`
	//FLVPseudoStreaming         *FLVPseudoStreaming           `json:"flvPseudoStreaming,omitempty"`
	//TimePseudoStreaming        *TimePseudoStreaming          `json:"timePseudoStreaming,omitempty"`
	//ResponseHeader             *ResponseHeader               `json:"responseHeader,omitempty"`
	//RedirectExceptions         *RedirectExceptions           `json:"redirectExceptions,omitempty"`
	//RedirectMappings           []*RedirectMappings           `json:"redirectMappings,omitempty"`
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
