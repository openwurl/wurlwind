package models

/*
	Int pointers are used so that we can have null ints and aren't erroneously sending 0's
	to the CDN configuration. 0 is a valid TTL, *int allows us to send 0's only when defined
	otherwise omitempty will leave out Null *int's
*/

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
	Enabled                        bool   `json:"enabled" tf:"enabled"`
	ExpirePolicy                   string `json:"expirePolicy" validate:"oneof=CACHE_CONTROL INGEST LAST_MODIFY NEVER_EXPIRE DO_NOT_CACHE" tf:"expire_policy"`
	ExpireSeconds                  *int   `json:"expireSeconds,omitempty" tf:"expire_seconds"`
	ForceBypassCache               bool   `json:"forceBypassCache,omitempty" tf:"force_bypass_cache"`
	HonorMustRevalidate            bool   `json:"honorMustRevalidate,omitempty" tf:"honor_must_revalidate"`
	HonorNoCache                   bool   `json:"honorNoCache,omitempty" tf:"honor_no_cache"`
	HonorNoStore                   bool   `json:"honorNoStore,omitempty" tf:"honor_no_store"`
	HonorPrivate                   bool   `json:"honorPrivate,omitempty" tf:"honor_private"`
	HonorSMaxAge                   bool   `json:"honorSMaxAge,omitempty" tf:"honor_smax_age"`
	HTTPHeaders                    string `json:"httpHeaders,omitempty" tf:"http_headers"` // string list
	MustRevalidateToNoCache        bool   `json:"mustRevalidateToNoCache,omitempty" tf:"must_revalidate_to_no_cache"`
	NoCacheBehavior                string `json:"noCacheBehavior,omitempty" tf:"no_cache_behavior"`
	UpdateHTTPHeadersOn304Response bool   `json:"updateHttpHeadersOn304Response,omitempty" tf:"update_http_headers_on_304_response"`
	DefaultCacheBehavior           string `json:"defaultCacheBehavior,omitempty" tf:"default_cache_behavior"` // Default behaviour when the policy is "Cache Control" and the "Cache-Control" header is missing. ttl & ...?
	MaxAgeZeroToNoCache            bool   `json:"maxAgeZeroToNoCache,omitempty" tf:"max_age_zero_to_no_cache"`
	BypassCacheIdentifier          string `json:"bypassCacheIdentifier,omitempty" tf:"bypass_cache_identifier"` // no-cache only
	ContentTypeFilter              string `json:"contentTypeFilter,omitempty" tf:"content_type_filter"`         // string list
	HeaderFilter                   string `json:"headerFilter,omitempty" tf:"header_filter"`                    // string list
	MethodFilter                   string `json:"methodFilter,omitempty" tf:"method_filter"`                    // string list
	PathFilter                     string `json:"pathFilter,omitempty" tf:"path_filter"`                        // string list
	StatusCodeMatch                string `json:"statusCodeMatch,omitempty" tf:"status_code_match"`             // string list
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

// GzipOriginPull ...
type GzipOriginPull struct {
	Enabled bool `json:"enabled"`
}

/**********************
Request & Response Modifications
*/

// TODO: This is MVP fields for modifications, however there are more that need implemented

// OriginRequestModification ...
type OriginRequestModification struct {
	Enabled     bool   `json:"enabled,omitempty"`
	AddHeaders  string `json:"addHeaders,omitempty"`
	FlowControl string `json:"flowControl,omitempty"`
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
	Enabled     bool   `json:"enabled,omitempty"`
	AddHeaders  string `json:"addHeaders,omitempty"`
	FlowControl string `json:"flowControl,omitempty"`
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
	Enabled     bool   `json:"enabled,omitempty"`
	AddHeaders  string `json:"addHeaders,omitempty"`
	FlowControl string `json:"flowControl,omitempty"`
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
	Enabled     bool   `json:"enabled,omitempty"`
	AddHeaders  string `json:"addHeaders,omitempty"`
	FlowControl string `json:"flowControl,omitempty"`
}

// Map converts the struct to a terraform consumable map
func (o *ClientRequestModification) Map() map[string]interface{} {
	mod := make(map[string]interface{})
	mod["enabled"] = o.Enabled
	mod["add_headers"] = o.AddHeaders
	mod["flow_control"] = o.FlowControl
	return mod
}

/**********************
Delivery Fields
*/

// Compression GZIP mime configuration
type Compression struct {
	Enabled bool   `json:"enabled,omitempty" tf:"enabled"`
	GZIP    string `json:"gzip,omitempty" tf:"gzip"`
	Level   int    `json:"level,string,omitempty" tf:"level"`
	Mime    string `json:"mime,omitempty" tf:"mime"`
}

// Map returns a terraform-consumable map of the compression struct
func (c *Compression) Map() map[string]interface{} {
	cm := make(map[string]interface{})
	cm["enabled"] = c.Enabled
	cm["gzip"] = c.GZIP
	cm["level"] = c.Level
	cm["mime"] = c.Mime
	return cm
}

// StaticHeader Headers to arbitrarily add
type StaticHeader struct {
	Enabled                  bool   `json:"enabled,omitempty" tf:"enabled"`
	HTTP                     string `json:"http,omitempty" tf:"http"`
	OriginPull               string `json:"originPull,omitempty" tf:"origin_pull"`
	ClientRequest            string `json:"clientRequest,omitempty" tf:"client_request"`
	MethodFilter             string `json:"methodFilter,omitempty" tf:"method_filter"` // comma delimited
	PathFilter               string `json:"pathFilter,omitempty" tf:"path_filter"`     // comma delimited
	HeaderFilter             string `json:"headerFilter,omitempty" tf:"header_filter"` // comma delimited
	ClientResponseCodeFilter string `json:"clientResponseCodeFilter,omitempty" tf:"client_response_code_filter"`
}

// Map returns a terraform-consumable map of the compression struct
func (s *StaticHeader) Map() map[string]interface{} {
	shm := make(map[string]interface{})
	shm["enabled"] = s.Enabled
	shm["origin_pull"] = s.OriginPull
	shm["client_request"] = s.ClientRequest
	shm["http"] = s.HTTP
	return shm
}

// HTTPMethods configures HTTP methods allowed
type HTTPMethods struct {
	Enabled  bool   `json:"enabled"`
	PassThru string `json:"passThru"`
}

// Map returns a terraform-consumable map of the compression struct
func (h *HTTPMethods) Map() map[string]interface{} {
	hmm := make(map[string]interface{})
	hmm["enabled"] = h.Enabled
	hmm["passthru"] = h.PassThru
	return hmm
}

// CustomMimeType ordered []CustomMimeType
type CustomMimeType struct {
	Enabled      bool   `json:"enabled,omitempty"`
	Code         string `json:"code,omitempty"`         // comma delimited
	ExtensionMap string `json:"extensionMap,omitempty"` // comma delimited
	MethodFilter string `json:"methodFilter,omitempty"` // comma delimited
	PathFilter   string `json:"pathFilter,omitempty"`   // comma delimited
	HeaderFilter string `json:"headerFilter,omitempty"` // comma delimited
}

// Map returns a terraform-consumable map of the custom mime type struct
func (c *CustomMimeType) Map() map[string]interface{} {
	cmt := make(map[string]interface{})
	cmt["enabled"] = c.Enabled
	cmt["code"] = c.Code
	cmt["extension_map"] = c.ExtensionMap
	cmt["method_filter"] = c.MethodFilter
	cmt["path_filter"] = c.PathFilter
	cmt["header_filter"] = c.HeaderFilter
	return cmt
}

// ContentDispositionByHeader ordered []ContentDispositionByHeader Controls the Content-Disposition header on the
// responses from the Origin using a pattern matched against the value of any
//HTTP header present in an end-user's request for content
type ContentDispositionByHeader struct {
	Enabled              bool   `json:"enabled,omitempty" tf:"enabled"`
	HeaderFieldName      string `json:"headerFieldName,omitempty" tf:"header_field_name"`
	HeaderValueMatch     string `json:"headerValueMatch,omitempty" tf:"header_value_match"` // comma delimited
	DefaultType          string `json:"defaultType,omitempty" validate:"oneof=inline attachment" tf:"default_type"`
	OverrideOriginHeader bool   `json:"overrideOriginHeader,omitempty" tf:"override_origin_header"`
	MethodFilter         string `json:"methodFilter,omitempty" tf:"method_filter"` // comma delimited
	PathFilter           string `json:"pathFilter,omitempty" tf:"path_filter"`     // comma delimited
	HeaderFilter         string `json:"headerFilter,omitempty" tf:"header_filter"` // comma delimited
}

// TODO maps for all

// BandwidthLimit ...
type BandwidthLimit struct {
	Enabled bool   `json:"enabled,omitempty"`
	Rule    string `json:"rule,omitempty"`   // | delimited
	Values  string `json:"values,omitempty"` // ex. 1mbps
}

// BandwidthRateLimit ...
type BandwidthRateLimit struct {
	Enabled            bool   `json:"enabled,omitempty"`
	InitialBurstName   string `json:"initialBurstName,omitempty"`   // ex. ri=
	SustainedRateName  string `json:"sustainedRateName,omitempty"`  // ex. rs=
	InitialBurstUnits  string `json:"initialBurstUnits,omitempty"`  // ex. byte
	SustainedRateUnits string `json:"sustainedRateUnits,omitempty"` // ex. kilobit
}

// DynamicCacheRule ordered []DynamicCacheRule
type DynamicCacheRule struct {
	Enabled      bool   `json:"enabled,omitempty"`
	MethodFilter string `json:"methodFilter,omitempty"` // comma delimited
	PathFilter   string `json:"pathFilter,omitempty"`   // comma delimited
	HeaderFilter string `json:"headerFilter,omitempty"` // comma delimited
	StatusCode   int    `json:"statusCode,omitempty"`
	Headers      string `json:"headers,omitempty"` // comma delimited
}

// FLVPseudoStreaming ...
type FLVPseudoStreaming struct {
	Enabled                     bool   `json:"enabled,omitempty"`
	JumpToByteInitialBytesParam string `json:"jumpToByteInitialBytesParam,omitempty"` // ex. ib
	JumpToByteStartOffsetParam  string `json:"jumpToByteStartOffsetParam,omitempty"`  // ex. fs
}

// TimePseudoStreaming ...
type TimePseudoStreaming struct {
	Enabled              bool   `json:"enabled,omitempty"`
	JumpToTimeStartParam string `json:"jumpToTimeStartParam,omitempty"` // ex. start
	JumpToTimeEndParam   string `json:"jumpToTimeEndParam,omitempty"`   // ex. end
}

// ResponseHeader ...
type ResponseHeader struct {
	Enabled     bool   `json:"enabled,omitempty"`
	HTTP        string `json:"http,omitempty"`
	EnabledETAg bool   `json:"enabledETag,omitempty"`
}

// RedirectExceptions ...
type RedirectExceptions struct {
	Enabled           bool   `json:"enabled,omitempty"`
	RedirectAgentCode string `json:"redirectAgentCode,omitempty"`
}

// RedirectMappings ordered []RedirectMappings
type RedirectMappings struct {
	Enabled          bool   `json:"enabled,omitempty"`
	Code             int    `json:"code,omitempty"`
	RedirectURL      string `json:"redirectURL,omitempty"`
	ReplacementToken string `json:"replacementToken,omitempty"`
}
