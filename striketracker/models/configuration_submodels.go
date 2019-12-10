package models

/**********************
Hostname
*/

// ScopeHostname is a single field that appears in a slice
// and contains the hostnames attached to a configuration scope
type ScopeHostname struct {
	Domain string `json:"domain"`
	Root   bool   `json:"root,omitempty"`
}

// ScopeHostnameFromInterfaceSlice returns a slice of ScopeHostname's from a slice of strings
func ScopeHostnameFromInterfaceSlice(hostnames []interface{}) []*ScopeHostname {
	shSlice := make([]*ScopeHostname, len(hostnames)-1)
	for _, hostname := range hostnames {
		thisHostname := &ScopeHostname{
			Domain: hostname.(string),
		}
		shSlice = append(shSlice, thisHostname)
	}
	return shSlice
}

/**********************
Origin Pull Cache Extension
*/

// OriginPullCacheExtension encapsulates stale cache extension settings
type OriginPullCacheExtension struct {
	Enabled                         bool `json:"enabled,omitempty"`
	ExpiredCacheExtension           *int `json:"expiredCacheExtension" validate:"required"`
	OriginUnreachableCacheExtension *int `json:"originUnreachableCacheExtension,omitempty"`
}

/**********************
Origin Pull host
*/

// OriginPullHost contains the origin IDs and path for a scope
type OriginPullHost struct {
	Primary   int    `json:"primary,omitempty"`
	Secondary int    `json:"secondary,omitempty"`
	Path      string `json:"path,omitempty"`
}

/**********************
Origin Pull Policy
*/

// OriginPullPolicy encapsulates origib pull policy settings
type OriginPullPolicy struct {
	Enabled                        bool   `json:"enabled"`
	ExpirePolicy                   string `json:"expirePolicy" validate:"oneof=CACHE_CONTROL INGEST LAST_MODIFY NEVER_EXPIRE DO_NOT_CACHE"`
	ExpireSeconds                  *int   `json:"expireSeconds"`
	ForceBypassCache               bool   `json:"forceBypassCache"`
	HonorMustRevalidate            bool   `json:"honorMustRevalidate"`
	HonorNoCache                   bool   `json:"honorNoCache"`
	HonorNoStore                   bool   `json:"honorNoStore"`
	HonorPrivate                   bool   `json:"honorPrivate"`
	HonorSMaxAge                   bool   `json:"honorSMaxAge"`
	HTTPHeaders                    string `json:"httpHeaders"` // string list
	MustRevalidateToNoCache        bool   `json:"mustRevalidateToNoCache"`
	NoCacheBehavior                string `json:"noCacheBehavior"`
	UpdateHTTPHeadersOn304Response bool   `json:"updateHttpHeadersOn304Response"`
	DefaultCacheBehavior           string `json:"defaultCacheBehavior"` // Default behaviour when the policy is "Cache Control" and the "Cache-Control" header is missing. ttl & ...?
	MaxAgeZeroToNoCache            bool   `json:"maxAgeZeroToNoCache"`
	BypassCacheIdentifier          string `json:"bypassCacheIdentifier"` // no-cache only
	ContentTypeFilter              string `json:"contentTypeFilter"`     // string list
	HeaderFilter                   string `json:"headerFilter"`          // string list
	MethodFilter                   string `json:"methodFilter"`          // string list
	PathFilter                     string `json:"pathFilter"`            // string list
	StatusCodeMatch                string `json:"statusCodeMatch"`       // string list
}

// NewOriginPullPolicyFromState returns a configured origin pull policy from a state index
func NewOriginPullPolicyFromState(state map[string]interface{}) *OriginPullPolicy {

	expireSeconds := state["expire_seconds"].(int)
	return &OriginPullPolicy{
		Enabled:                        state["enabled"].(bool),
		ExpirePolicy:                   state["expire_policy"].(string),
		ExpireSeconds:                  &expireSeconds,
		ForceBypassCache:               state["force_bypass_cache"].(bool),
		HonorMustRevalidate:            state["honor_must_revalidate"].(bool),
		HonorNoCache:                   state["honor_must_revalidate"].(bool),
		HonorNoStore:                   state["honor_no_store"].(bool),
		HonorPrivate:                   state["honor_private"].(bool),
		HonorSMaxAge:                   state["honor_smax_age"].(bool),
		HTTPHeaders:                    state["http_headers"].(string),
		MustRevalidateToNoCache:        state["must_revalidate_to_no_cache"].(bool),
		NoCacheBehavior:                state["no_cache_behavior"].(string),
		UpdateHTTPHeadersOn304Response: state["update_http_headers_on_304_response"].(bool),
		DefaultCacheBehavior:           state["default_cache_behavior"].(string),
		MaxAgeZeroToNoCache:            state["max_age_zero_to_no_cache"].(bool),
		BypassCacheIdentifier:          state["bypass_cache_identifier"].(string),
		ContentTypeFilter:              state["content_type_filter"].(string),
		HeaderFilter:                   state["header_filter"].(string),
		MethodFilter:                   state["method_filter"].(string),
		PathFilter:                     state["path_filter"].(string),
		StatusCodeMatch:                state["status_code_match"].(string),
	}
}

/**********************
Request & Response Modifications
*/

// TODO: This is MVP fields for modifications, however there are more that need implemented

// OriginRequestModification ...
type OriginRequestModification struct {
	Enabled     bool   `json:"enabled"`
	AddHeaders  string `json:"addHeaders"`
	FlowControl string `json:"flowControl"`
}

// Map converts the struct to a terraform consumable map
func (o *OriginRequestModification) Map() map[string]interface{} {
	mod := make(map[string]interface{})
	mod["enabled"] = o.Enabled
	mod["add_headers"] = o.AddHeaders
	mod["flow_control"] = o.FlowControl
	return mod
}

// OriginResponseModification ...
type OriginResponseModification struct {
	Enabled     bool   `json:"enabled"`
	AddHeaders  string `json:"addHeaders"`
	FlowControl string `json:"flowControl"`
}

// Map converts the struct to a terraform consumable map
func (o *OriginResponseModification) Map() map[string]interface{} {
	mod := make(map[string]interface{})
	mod["enabled"] = o.Enabled
	mod["add_headers"] = o.AddHeaders
	mod["flow_control"] = o.FlowControl
	return mod
}

// ClientResponseModification ...
type ClientResponseModification struct {
	Enabled     bool   `json:"enabled"`
	AddHeaders  string `json:"addHeaders"`
	FlowControl string `json:"flowControl"`
}

// Map converts the struct to a terraform consumable map
func (o *ClientResponseModification) Map() map[string]interface{} {
	mod := make(map[string]interface{})
	mod["enabled"] = o.Enabled
	mod["add_headers"] = o.AddHeaders
	mod["flow_control"] = o.FlowControl
	return mod
}

// ClientRequestModification ...
type ClientRequestModification struct {
	Enabled     bool   `json:"enabled"`
	AddHeaders  string `json:"addHeaders"`
	FlowControl string `json:"flowControl"`
}

// Map converts the struct to a terraform consumable map
func (o *ClientRequestModification) Map() map[string]interface{} {
	mod := make(map[string]interface{})
	mod["enabled"] = o.Enabled
	mod["add_headers"] = o.AddHeaders
	mod["flow_control"] = o.FlowControl
	return mod
}
