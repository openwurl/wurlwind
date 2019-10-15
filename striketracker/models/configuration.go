package models

import (
	"fmt"
	"strings"

	"github.com/openwurl/wurlwind/pkg/utilities"
	"github.com/openwurl/wurlwind/pkg/validation"
	validator "gopkg.in/go-playground/validator.v9"
)

// ValidPullProtocols are for matching against user input on pull protocol
var ValidPullProtocols = []string{"http", "https", "match"}

// ValidExpirePolicies are for matching against user input on expire policies
var ValidExpirePolicies = []string{"CACHE_CONTROL", "INGEST", "LAST_MODIFY", "NEVER_EXPIRE", "DO_NOT_CACHE"}

// Configuration defines a high level scope configuration for a delivery hash
type Configuration struct {
	Response
	Hostname                    []*ConfigurationHostname      `json:"hostname"`
	OriginPullLogs              *OriginPullLogs               `json:"originPullLogs"`
	OriginPullProtocol          *OriginPullProtocol           `json:"originPullProtocol"`
	OriginPullPolicy            []*OriginPullPolicy           `json:"originPullPolicy"`
	FileSegmentation            *FileSegmentation             `json:"fileSegmentation"`
	GzipOriginPull              *GzipOriginPull               `json:"gzipOriginPull"`
	OriginPersistentConnections *OriginPersistentConnections  `json:"originPersistentConnections"`
	OriginPull                  *OriginPull                   `json:"originPull"`
	CacheControl                []*CacheControl               `json:"cacheControl"`
	CacheKeyModification        *CacheKeyModification         `json:"cacheKeyModification"`
	Compression                 *Compression                  `json:"compression"`
	StaticHeader                []*StaticHeader               `json:"staticHeader"`
	HTTPMethods                 *HTTPMethods                  `json:"httpMethods"`
	AccessLogs                  *AccessLogs                   `json:"accessLogs"`
	OriginPullHost              *OriginPullHost               `json:"originPullHost"`
	OriginRequestModification   []*OriginRequestModification  `json:"originRequestModification,omitempty"`
	OriginResponseModification  []*OriginResponseModification `json:"originResponseModification,omitempty"`
	ClientRequestModification   []*ClientRequestModification  `json:"clientRequestModification,omitempty"`
	ClientResponseModification  []*ClientResponseModification `json:"clientResponseModification,omitempty"`
	*Scope                      `json:"scope"`
}

// Validate validates the struct data
func (c *Configuration) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(c); err != nil {
		return err
	}
	return nil
}

// BuildHostScopeInterface returns scope details
// for the host in a terraform compatible interface
func (c *Configuration) BuildHostScopeInterface() map[string]interface{} {
	scopeList := make(map[string]interface{})
	scopeList["id"] = string(c.ID)
	scopeList["platform"] = c.Platform
	scopeList["path"] = c.Path
	return scopeList
}

// BuildOriginPullPoliciesInterface returns a tf-compatible interface from the model
func (c *Configuration) BuildOriginPullPoliciesInterface() []interface{} {
	policies := c.OriginPullPolicy
	var policyList []interface{}
	for _, policy := range policies {
		thisPolicy := make(map[string]interface{})
		thisPolicy["enabled"] = policy.Enabled
		thisPolicy["expire_seconds"] = policy.ExpireSeconds
		thisPolicy["force_bypass_cache"] = policy.ForceBypassCache
		thisPolicy["honor_must_revalidate"] = policy.HonorMustRevalidate
		thisPolicy["honor_no_cache"] = policy.HonorNoCache
		thisPolicy["honor_no_store"] = policy.HonorNoStore
		thisPolicy["honor_private"] = policy.HonorPrivate
		thisPolicy["honor_smax_age"] = policy.HonorSMaxAge
		thisPolicy["http_headers"] = policy.HTTPHeaders
		thisPolicy["must_revalidate_to_no_cache"] = policy.MustRevalidateToNoCache
		thisPolicy["no_cache_behavior"] = policy.NoCacheBehavior
		thisPolicy["update_http_headers_on_304_response"] = policy.UpdateHTTPHeadersOn304Response
		thisPolicy["default_cache_behavior"] = policy.DefaultCacheBehavior
		thisPolicy["max_age_zero_to_no_cache"] = policy.MaxAgeZeroToNoCache
		thisPolicy["content_type_filter"] = policy.ContentTypeFilter
		thisPolicy["header_filter"] = policy.HeaderFilter
		thisPolicy["method_filter"] = policy.MethodFilter
		thisPolicy["path_filter"] = policy.PathFilter
		policyList = append(policyList, thisPolicy)
	}
	return policyList
}

// ClientResponseMap returns a tf interface slice of client request mods
func (c *Configuration) ClientResponseMap() []interface{} {
	thisMap := make([]interface{}, len(c.ClientResponseModification))
	for _, mod := range c.ClientResponseModification {
		thisMap = append(thisMap, mod.AsMap())
	}
	return thisMap
}

// ClientRequestMap returns a tf interface slice of client request mods
func (c *Configuration) ClientRequestMap() []interface{} {
	thisMap := make([]interface{}, len(c.ClientRequestModification))
	for _, mod := range c.ClientRequestModification {
		thisMap = append(thisMap, mod.AsMap())
	}
	return thisMap
}

// OriginRequestMap returns a tf interface slice of client request mods
func (c *Configuration) OriginRequestMap() []interface{} {
	thisMap := make([]interface{}, len(c.OriginRequestModification))
	for _, mod := range c.OriginRequestModification {
		thisMap = append(thisMap, mod.AsMap())
	}
	return thisMap
}

// OriginResponseMap returns a tf interface slice of client request mods
func (c *Configuration) OriginResponseMap() []interface{} {
	thisMap := make([]interface{}, len(c.OriginResponseModification))
	for _, mod := range c.OriginResponseModification {
		thisMap = append(thisMap, mod.AsMap())
	}
	return thisMap
}

// HostnamesAsList returns an list of strings in an interface for tf state
func (c *Configuration) HostnamesAsList() []interface{} {
	ret := make([]interface{}, len(c.Hostname))
	for _, host := range c.Hostname {
		ret = append(ret, host.Domain)
	}
	return ret
}

// HostnamesAsStringSlice returns a slice of strings for tf state
func (c *Configuration) HostnamesAsStringSlice() []string {
	hostnames := make([]string, len(c.Hostname))
	for _, hostname := range c.Hostname {
		if hostname.Domain == "" {
			// Skip blank fields
			continue
		}
		hostnames = append(hostnames, hostname.Domain)
	}
	return hostnames
}

// ActionableHostnamesAsStringSlice returns a pared down slice of strings
// only containing those set by the user
func (c *Configuration) ActionableHostnamesAsStringSlice() []string {
	hostnames := make([]string, len(c.Hostname))
	for _, hostname := range c.Hostname {
		if hostname.Domain == "" || strings.Contains(hostname.Domain, "hwcdn.net") {
			// Skip blank fields
			continue
		}
		hostnames = append(hostnames, hostname.Domain)
	}
	return hostnames
}

// BuildOriginInterface returns a tf state
// compatible reflection of origin pull host and other details
func (c *Configuration) BuildOriginInterface() map[string]interface{} {
	originMap := make(map[string]interface{})
	if c.OriginPullHost != nil {
		originMap["primary"] = c.OriginPullHost.Primary
		originMap["secondary"] = c.OriginPullHost.Secondary
		originMap["path"] = c.OriginPullHost.Path
	}
	originMap["origin_pull_protocol"] = c.OriginPullProtocol.Protocol
	originMap["redirect_action"] = c.OriginPull.RedirectAction

	return originMap
}

// ConfigurationCreate because POST a new config is a unicorn due to bad API design
type ConfigurationCreate struct {
	Response
	Hostname                    []*ConfigurationHostname     `json:"hostname"`
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
	Name                        string                       `json:"name"`
	Platform                    string                       `json:"platform" validate:"required"`
	Path                        string                       `json:"path" validate:"required"`
	ID                          int                          `json:"id"`
}

// Validate validates the struct data
func (c *ConfigurationCreate) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(c); err != nil {
		return err
	}
	return nil
}

// NewDefaultConfiguration returns a baseline configuration to be modified with defaults
func NewDefaultConfiguration() *Configuration {
	c := &Configuration{
		OriginPullLogs: &OriginPullLogs{
			Enabled: true,
		},
		OriginPullProtocol: &OriginPullProtocol{
			Protocol: "https",
		},
		FileSegmentation: &FileSegmentation{
			Enabled: true,
		},
		GzipOriginPull: &GzipOriginPull{
			Enabled: true,
		},
		OriginPersistentConnections: &OriginPersistentConnections{
			Enabled: false,
		},
		OriginPull: &OriginPull{
			RedirectAction: "proxy",
		},
		CacheKeyModification: &CacheKeyModification{
			NormalizeKeyPathToLowerCase: true,
		},
		Compression: &Compression{
			GZIP: "txt,js,htm,html,css",
			Mime: "test/*",
		},
		HTTPMethods: &HTTPMethods{
			PassThru: "*",
		},
		AccessLogs: &AccessLogs{
			Enabled: true,
		},
		OriginPullHost: &OriginPullHost{},
	}
	c.OriginPullPolicy = append(c.OriginPullPolicy, &OriginPullPolicy{
		ExpirePolicy:                   "CACHE_CONTROL",
		ExpireSeconds:                  31536000,
		ForceBypassCache:               false,
		HonorMustRevalidate:            true,
		HonorNoCache:                   true,
		HonorPrivate:                   true,
		HonorSMaxAge:                   true,
		HTTPHeaders:                    "*",
		MustRevalidateToNoCache:        true,
		NoCacheBehavior:                "spec",
		UpdateHTTPHeadersOn304Response: true,
	})
	c.CacheControl = append(c.CacheControl, &CacheControl{
		MaxAge:            31536000,
		SynchronizeMaxAge: true,
	})
	c.StaticHeader = append(c.StaticHeader, &StaticHeader{
		HTTP:       "Access-Control-Allow-Origin: *",
		OriginPull: "Host: %client.request.host%",
	})
	return c
}

/*
	Configuration Modification
*/

// SetOriginPullLogs enables or disables origin pull logging
func (c *Configuration) SetOriginPullLogs(enabled bool) {
	c.OriginPullLogs.Enabled = enabled
}

// GetOriginPullLogs returns the origin pull log enabled setting
func (c *Configuration) GetOriginPullLogs() bool {
	return c.OriginPullLogs.Enabled
}

// SetOriginPullProtocol sets the origin pull protocol to the one given
func (c *Configuration) SetOriginPullProtocol(protocol string) error {
	if !utilities.SliceContainsString(protocol, ValidPullProtocols) {
		return fmt.Errorf("%s is not a valid protocol. Must be one of (%s)", protocol, strings.Join(ValidPullProtocols, ","))
	}
	c.OriginPullProtocol.Protocol = protocol
	return nil
}

// SetFileSegmentation

// SetGzipOriginPull

// SetOriginPersistentConnections

// SetOriginPull

/*
	Sub structures
*/

// OriginRequestModification ...
type OriginRequestModification struct {
	Enabled     bool   `json:"enabled"`
	AddHeaders  string `json:"addHeaders"`
	FlowControl string `json:"flowControl"`
}

// AsMap converts the struct to a terraform consumable map
func (o *OriginRequestModification) AsMap() map[string]interface{} {
	mod := make(map[string]interface{})
	mod["enabled"] = o.Enabled
	mod["add_headers"] = o.AddHeaders
	mod["flow_control"] = o.FlowControl
	return mod
}

// BuildOriginRequestModification ...
func BuildOriginRequestModification(tfSchema []interface{}) []*OriginRequestModification {
	modList := []*OriginRequestModification{}

	// extract
	for _, mod := range tfSchema {
		thisModRaw := mod.(map[string]interface{})
		thisMod := &OriginRequestModification{
			Enabled:     thisModRaw["enabled"].(bool),
			AddHeaders:  thisModRaw["add_headers"].(string),
			FlowControl: thisModRaw["flow_control"].(string),
		}
		modList = append(modList, thisMod)
	}
	return modList
}

// OriginResponseModification ...
type OriginResponseModification struct {
	Enabled     bool   `json:"enabled"`
	AddHeaders  string `json:"addHeaders"`
	FlowControl string `json:"flowControl"`
}

// AsMap converts the struct to a terraform consumable map
func (o *OriginResponseModification) AsMap() map[string]interface{} {
	mod := make(map[string]interface{})
	mod["enabled"] = o.Enabled
	mod["add_headers"] = o.AddHeaders
	mod["flow_control"] = o.FlowControl
	return mod
}

// BuildOriginResponseModification ...
func BuildOriginResponseModification(tfSchema []interface{}) []*OriginResponseModification {
	modList := []*OriginResponseModification{}

	// extract
	for _, mod := range tfSchema {
		thisModRaw := mod.(map[string]interface{})
		thisMod := &OriginResponseModification{
			Enabled:     thisModRaw["enabled"].(bool),
			AddHeaders:  thisModRaw["add_headers"].(string),
			FlowControl: thisModRaw["flow_control"].(string),
		}
		modList = append(modList, thisMod)
	}
	return modList
}

// ClientResponseModification ...
type ClientResponseModification struct {
	Enabled     bool   `json:"enabled"`
	AddHeaders  string `json:"addHeaders"`
	FlowControl string `json:"flowControl"`
}

// AsMap converts the struct to a terraform consumable map
func (o *ClientResponseModification) AsMap() map[string]interface{} {
	mod := make(map[string]interface{})
	mod["enabled"] = o.Enabled
	mod["add_headers"] = o.AddHeaders
	mod["flow_control"] = o.FlowControl
	return mod
}

// BuildClientResponseModification ...
func BuildClientResponseModification(tfSchema []interface{}) []*ClientResponseModification {
	modList := []*ClientResponseModification{}

	// extract
	for _, mod := range tfSchema {
		thisModRaw := mod.(map[string]interface{})
		thisMod := &ClientResponseModification{
			Enabled:     thisModRaw["enabled"].(bool),
			AddHeaders:  thisModRaw["add_headers"].(string),
			FlowControl: thisModRaw["flow_control"].(string),
		}
		modList = append(modList, thisMod)
	}
	return modList
}

// ClientRequestModification ...
type ClientRequestModification struct {
	Enabled     bool   `json:"enabled"`
	AddHeaders  string `json:"addHeaders"`
	FlowControl string `json:"flowControl"`
}

// AsMap converts the struct to a terraform consumable map
func (o *ClientRequestModification) AsMap() map[string]interface{} {
	mod := make(map[string]interface{})
	mod["enabled"] = o.Enabled
	mod["add_headers"] = o.AddHeaders
	mod["flow_control"] = o.FlowControl
	return mod
}

// BuildClientRequestModification ...
func BuildClientRequestModification(tfSchema []interface{}) []*ClientRequestModification {
	modList := []*ClientRequestModification{}

	// extract
	for _, mod := range tfSchema {
		thisModRaw := mod.(map[string]interface{})
		thisMod := &ClientRequestModification{
			Enabled:     thisModRaw["enabled"].(bool),
			AddHeaders:  thisModRaw["add_headers"].(string),
			FlowControl: thisModRaw["flow_control"].(string),
		}
		modList = append(modList, thisMod)
	}
	return modList
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
	Enabled                        bool   `json:"enabled"`
	ExpirePolicy                   string `json:"expirePolicy" validate:"oneof=CACHE_CONTROL INGEST LAST_MODIFY NEVER_EXPIRE DO_NOT_CACHE"`
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
	DefaultCacheBehavior           string `json:"defaultCacheBehavior"` // Default behaviour when the policy is "Cache Control" and the "Cache-Control" header is missing.
	MaxAgeZeroToNoCache            bool   `json:"maxAgeZeroToNoCache"`
	ContentTypeFilter              string `json:"contentTypeFilter"`
	HeaderFilter                   string `json:"headerFilter"`
	MethodFilter                   string `json:"methodFilter"`
	PathFilter                     string `json:"pathFilter"`
}

// AsMap returns the object as a TF-consumable map
func (o *OriginPullPolicy) AsMap() map[string]interface{} {
	thisMap := make(map[string]interface{})
	// TODO: this
	return thisMap
}

// NewOriginPullPolicyListFromInterface returns a slice of policies from a terraform interface
func NewOriginPullPolicyListFromInterface(terraformPullPolicyList *[]interface{}) []*OriginPullPolicy {
	policylist := []*OriginPullPolicy{}

	for _, policy := range *terraformPullPolicyList {
		thisMap := policy.(map[string]interface{})

		newPolicy := &OriginPullPolicy{
			Enabled:                        thisMap["enabled"].(bool),
			ExpirePolicy:                   thisMap["expire_policy"].(string),
			ExpireSeconds:                  thisMap["expire_seconds"].(int),
			ForceBypassCache:               thisMap["force_bypass_cache"].(bool),
			HonorMustRevalidate:            thisMap["honor_must_revalidate"].(bool),
			HonorNoCache:                   thisMap["honor_no_cache"].(bool),
			HonorNoStore:                   thisMap["honor_no_store"].(bool),
			HonorPrivate:                   thisMap["honor_private"].(bool),
			HonorSMaxAge:                   thisMap["honor_smax_age"].(bool),
			HTTPHeaders:                    thisMap["http_headers"].(string),
			MustRevalidateToNoCache:        thisMap["must_revalidate_to_no_cache"].(bool),
			NoCacheBehavior:                thisMap["no_cache_behavior"].(string),
			UpdateHTTPHeadersOn304Response: thisMap["update_http_headers_on_304_response"].(bool),
			DefaultCacheBehavior:           thisMap["default_cache_behavior"].(string),
			MaxAgeZeroToNoCache:            thisMap["max_age_zero_to_no_cache"].(bool),
			ContentTypeFilter:              thisMap["content_type_filter"].(string),
			HeaderFilter:                   thisMap["header_filter"].(string),
			MethodFilter:                   thisMap["method_filter"].(string),
			PathFilter:                     thisMap["path_filter"].(string),
		}
		policylist = append(policylist, newPolicy)
	}

	return policylist
}

// ConfigurationHostname ...
type ConfigurationHostname struct {
	Domain string `json:"domain"`
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
	GZIP  string `json:"gzip"`
	Level int    `json:"level,string"`
	Mime  string `json:"mime"`
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
	Primary   int    `json:"primary"`
	Secondary int    `json:"secondary"`
	Path      string `json:"path"`
}

// BuildOriginInterface returns a tf compatible map for state
func (o *OriginPullHost) BuildOriginInterface() map[string]interface{} {
	return nil
}

// ConfigurationScope is the scope name
type ConfigurationScope struct {
	Name string `json:"name"`
}
