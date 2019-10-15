package models

import "strings"

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

// AttachLogState adds log state to the existing configuration from terraform state
func (c *Configuration) AttachLogState(state map[string]interface{}) {
	if al := state["access_logs"]; al != nil {
		c.AccessLogs = &AccessLogs{
			Enabled: al.(bool),
		}
	}

	if opl := state["origin_pull_logs"]; opl != nil {
		c.OriginPullLogs = &OriginPullLogs{
			Enabled: opl.(bool),
		}
	}
}

// GetLogState returns a tf-compatible map of access and origin_pull logs
func (c *Configuration) GetLogState() map[string]interface{} {
	ls := make(map[string]interface{})
	ls["access_logs"] = c.AccessLogs.Enabled
	ls["origin_pull_logs"] = c.OriginPullLogs.Enabled
	return ls
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
	originMap["gzip"] = c.GzipOriginPull.Enabled
	originMap["persistent_connections"] = c.OriginPersistentConnections.Enabled

	return originMap
}

// BuildCacheControlInterface returns a tf state compatible
// interface of cache controls
func (c *Configuration) BuildCacheControlInterface() *[]interface{} {
	ccIface := make([]interface{}, len(c.CacheControl))
	for _, cc := range c.CacheControl {
		thisCC := make(map[string]interface{})
		thisCC["enabled"] = cc.Enabled
		thisCC["must_revalidate"] = cc.MustRevalidate
		thisCC["max_age"] = cc.MaxAge
		thisCC["synchronize_max_age"] = cc.SynchronizeMaxAge
		thisCC["override"] = cc.Override
		ccIface = append(ccIface, thisCC)
	}

	return &ccIface
}
