package models

// Exported defaults for model data
const (
	DefaultCompressionEnabledValue = true
	DefaultCompressionLevelValue   = 1
	DefaultCompressionMimeValue    = "test/*"
	DefaultCompressionGZIPValue    = "txt,js,htm,html,css"
)

// Exported lists of valid settings for strict validation
var (
	ValidPullProtocols    = []string{"http", "https", "match"}
	ValidExpirePolicies   = []string{"CACHE_CONTROL", "INGEST", "LAST_MODIFY", "NEVER_EXPIRE", "DO_NOT_CACHE"}
	ValidRedirectActions  = []string{"proxy", "follow"}
	ValidNoCacheBehaviors = []string{"spec", "legacy"}
	ValidCacheBehaviors   = []string{"ttl", "bypass"}
)
