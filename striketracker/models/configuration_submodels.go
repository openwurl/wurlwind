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

/*
	Origin Pull Cache Extension
*/

// OriginPullCacheExtension encapsulates stale cache extension settings
type OriginPullCacheExtension struct {
	Enabled                         bool `json:"enabled,omitempty"`
	ExpiredCacheExtension           int  `json:"expiredCacheExtension,omitempty"`
	OriginUnreachableCacheExtension int  `json:"originUnreachableCacheExtension,omitempty"`
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
