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
	DefaultCacheBehavior           string `json:"defaultCacheBehavior"` // Default behaviour when the policy is "Cache Control" and the "Cache-Control" header is missing.
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
	return &OriginPullPolicy{
		Enabled: state["enabled"].(bool),
	}
}

/*
   "id": 538230093,
   "statusCodeMatch": "200,201",
   "expirePolicy": "CACHE_CONTROL",
   "expireSeconds": 1,
   "honorNoStore": true,
   "honorNoCache": true,
   "honorMustRevalidate": true,
   "noCacheBehavior": "spec",
   "maxAgeZeroToNoCache": true,
   "mustRevalidateToNoCache": true,
   "forceBypassCache": true,
   "httpHeaders": "Access-Control-Allow-Origin,x-test-thing",
   "honorPrivate": true,
   "honorSMaxAge": true,
   "updateHttpHeadersOn304Response": true,
   "defaultCacheBehavior": "ttl",
   "enabled": true,
   "methodFilter": "GET",
   "pathFilter": "*filter1*,filter2",
   "headerFilter": "*header_filter",
   "contentTypeFilter": "*"
*/
