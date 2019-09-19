package models

import (
	"github.com/openwurl/wurlwind/pkg/validation"
	validator "gopkg.in/go-playground/validator.v9"
)

// Configuration defines a high level scope configuration for a delivery hash
type Configuration struct {
	Response
	Hostname                    []string                     `json:"hostname"`
	OriginPullLogs              *OriginPullLogs              `json:"originPullLogs"`
	OriginPullProtocol          *OriginPullProtocol          `json:"originPullProtocol"`
	OriginPullPolicy            []*OriginPullPolicy          `json:"originPullPolicy"`
	FileSegmentation            *FileSegmentation            `json:"fileSegmentation"`
	GzipOriginPull              *GzipOriginPull              `json:"gzipOriginPull"`
	OriginPersistentConnections *OriginPersistentConnections `json:"originPersistentConnections"`
	OriginPull                  *OriginPull                  `json:"originPull"`
	CacheControl                []*CacheControl              `json:"cacheControl"`
	CacheKeyModification        *CacheKeyModification        `json:"cacheKeyModification"`
	Compression                 *Compression                 `json:"compression"`
	StaticHeader                []*StaticHeader              `json:"staticHeader"`
	HTTPMethods                 *HTTPMethods                 `json:"httpMethods"`
	AccessLogs                  *AccessLogs                  `json:"accessLogs"`
	OriginPullHost              *OriginPullHost              `json:"originPullHost"`
	Scope                       *ConfigurationScope          `json:"scope"`
}

// Validate validates the struct data
func (c *Configuration) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(c); err != nil {
		return err
	}
	return nil
}

// OriginPullLogs encapsulates origin pull log settings
type OriginPullLogs struct {
	Enabled bool `json:"enabled"`
}

// OriginPullProtocol encapsulates origin pull log settings
type OriginPullProtocol struct {
	Protocol string `json:"protocol"`
}

// OriginPullPolicy encapsulates origib pull policy settings
type OriginPullPolicy struct {
	ExpirePolicy                   string `json:"expirePolicy"`
	ExpireSeconds                  int    `json:"expireSeconds"`
	ForceBypassCache               bool   `json:"forceBypassCache"`
	HonorMustRevalidate            bool   `json:"honorMustRevalidate"`
	HonorNoCache                   bool   `json:"honorNoCache"`
	HonorNoStore                   bool   `json:"honorNoStore"`
	HonorPrivate                   bool   `json:"honorPrivate"`
	HonorSMaxAge                   bool   `json:"honorSMaxAge"`
	HTTPHeaders                    string `json:"httpHeaders"`
	MustRevalidateToNoCache        bool   `json:"mustRevalidateToNoCache"`
	NoCacheBehavior                string `json:"noCacheBehavior"`
	UpdateHTTPHeadersOn304Response bool   `json:"updateHttpHeadersOn304Response"`
}

// FileSegmentation ...
type FileSegmentation struct {
	Enabled bool `json:"enabled"`
}

// GzipOriginPull ...
type GzipOriginPull struct {
	Enabled bool `json:"enabled"`
}

// OriginPersistentConnections ...
type OriginPersistentConnections struct {
	Enabled bool `json:"enabled"`
}

// OriginPull ...
type OriginPull struct {
	RedirectAction string `json:"redirectAction"`
}

// CacheControl ...
type CacheControl struct {
	MaxAge            int  `json:"maxAge"`
	SynchronizeMaxAge bool `json:"synchronizeMaxAge"`
}

// CacheKeyModification ...
type CacheKeyModification struct {
	NormalizeKeyPathToLowerCase bool `json:"normalizeKeyPathToLowerCase"`
}

// Compression GZIP mime configuration
type Compression struct {
	GZIP string `json:"gzip"`
	Mime string `json:"mime"`
}

// StaticHeader Headers to arbitrarily add
type StaticHeader struct {
	HTTP       string `json:"http"`
	OriginPull string `json:"originPull"`
}

// HTTPMethods configures HTTP methods allowed
type HTTPMethods struct {
	PassThru string `json:"passThru"`
}

// AccessLogs defines whether or not access logging is enabled
type AccessLogs struct {
	Enabled bool `json:"enabled"`
}

// OriginPullHost contains the origin ID for this scope configuration
type OriginPullHost struct {
	Primary   int `json:"primary"`
	Secondary int `json:"secondary"`
}

// ConfigurationScope is the scope name
type ConfigurationScope struct {
	Name string `json:"name"`
}
