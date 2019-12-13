package models

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/openwurl/wurlwind/pkg/debug"
)

// =========
// Scope

// ScopeFromState updates the models scope details from state
func (c *Configuration) ScopeFromState(state map[string]interface{}) {
	c.Scope = &Scope{}
	debug.Log("scope-state", "PLATFORM: %v", state["platform"])
	if state["platform"] != nil {
		c.Platform = state["platform"].(string)
	}
	if state["path"] != nil {
		c.Path = state["path"].(string)
	}
	if state["name"] != nil {
		c.Name = state["name"].(string)
	}
}

// ScopeFromModel returns a scope state hash from the model
func (c *Configuration) ScopeFromModel() map[string]interface{} {
	scopeIface := make(map[string]interface{})
	scopeIface["platform"] = c.Platform
	scopeIface["path"] = c.Path
	scopeIface["name"] = c.Name
	return scopeIface
}

// =========
// Hostnames

// HostnamesFromState appends a hostname interface slice to the model
func (c *Configuration) HostnamesFromState(hostnames []interface{}) {
	hostnameList := ScopeHostnameFromInterfaceSlice(hostnames)
	c.Hostname = hostnameList
}

// HostnamesFromModel returns a slice interface of hostnames
func (c *Configuration) HostnamesFromModel() []interface{} {
	// TODO: correct the size of this make
	hostnameIF := make([]interface{}, 0)
	for _, hostname := range c.Hostname {
		hostnameIF = append(hostnameIF, hostname.Domain)
	}
	return hostnameIF
}

// =========
// Origin Host

// OriginHostFromState ...
func (c *Configuration) OriginHostFromState(state map[string]interface{}) {
	c.OriginPullHost = &OriginPullHost{}
	if state["primary"] != nil {
		c.OriginPullHost.Primary = state["primary"].(int)
	}
	if state["secondary"] != nil {
		c.OriginPullHost.Secondary = state["secondary"].(int)
	}
	if state["path"] != nil {
		c.OriginPullHost.Path = state["path"].(string)
	}
}

// OriginHostFromModel ...
func (c *Configuration) OriginHostFromModel() []interface{} {
	originHostSliceIface := []interface{}{}
	originHostIface := make(map[string]interface{})

	debug.Log("ORIGIN_HOST", "%v", spew.Sprintf("%v", c.OriginPullHost))

	if c.OriginPullHost != nil {
		originHostIface["primary"] = c.OriginPullHost.Primary
		originHostIface["secondary"] = c.OriginPullHost.Secondary
		originHostIface["path"] = c.OriginPullHost.Path
	} else {
		originHostIface["primary"] = nil
		originHostIface["secondary"] = nil
		originHostIface["path"] = nil
	}
	originHostSliceIface = append(originHostSliceIface, originHostIface)

	return originHostSliceIface
}

// =========
// Origin Pull Cache Extension / stale_cache_extension

// OriginPullCacheExtensionFromState ...
func (c *Configuration) OriginPullCacheExtensionFromState(state map[string]interface{}) {
	c.OriginPullCacheExtension = &OriginPullCacheExtension{}
	if state["enabled"] != nil {
		c.OriginPullCacheExtension.Enabled = state["enabled"].(bool)
	}
	if state["expired_cache_extension"] != nil {
		ece := state["expired_cache_extension"].(int)
		c.OriginPullCacheExtension.ExpiredCacheExtension = &ece
	}
	if state["origin_unreachable_cache_extension"] != nil {
		ouce := state["origin_unreachable_cache_extension"].(int)
		c.OriginPullCacheExtension.OriginUnreachableCacheExtension = &ouce
	}
}

// OriginPullCacheExtensionFromModel ...
func (c *Configuration) OriginPullCacheExtensionFromModel() []interface{} {
	originPullCacheExtensionSliceIface := []interface{}{}
	originPullCacheExtensionIface := make(map[string]interface{})

	if c.OriginPullCacheExtension != nil {
		originPullCacheExtensionIface["enabled"] = c.OriginPullCacheExtension.Enabled
		originPullCacheExtensionIface["expired_cache_extension"] = c.OriginPullCacheExtension.ExpiredCacheExtension
		originPullCacheExtensionIface["origin_unreachable_cache_extension"] = c.OriginPullCacheExtension.OriginUnreachableCacheExtension
	}

	originPullCacheExtensionSliceIface = append(originPullCacheExtensionSliceIface, originPullCacheExtensionIface)

	return originPullCacheExtensionSliceIface
}

// =========
// Origin Pull Policy / cache_policy

// OriginPullPolicyFromState ...
func (c *Configuration) OriginPullPolicyFromState(state []interface{}) error {
	orderedList := make([]interface{}, len(state))
	c.OriginPullPolicy = make([]*OriginPullPolicy, 0)

	// order the list by defined weight
	for _, policy := range state {
		policyCast := policy.(map[string]interface{})
		policyIndex := policyCast["weight"].(int)

		if orderedList[policyIndex] != nil {
			return fmt.Errorf("Weight %d used multiple times", policyIndex)
		}
		orderedList[policyIndex] = policyCast
	}

	for _, policy := range orderedList {
		thisPolicy := NewOriginPullPolicyFromState(policy.(map[string]interface{}))
		c.OriginPullPolicy = append(c.OriginPullPolicy, thisPolicy)
	}

	return nil

}

// OriginPullPolicyFromModel ...
func (c *Configuration) OriginPullPolicyFromModel() []interface{} {
	originPullPolicyIface := make([]interface{}, 0)

	for index, policy := range c.OriginPullPolicy {
		thisPolicy := make(map[string]interface{})
		thisPolicy["weight"] = index
		thisPolicy["enabled"] = policy.Enabled
		thisPolicy["expire_policy"] = policy.ExpirePolicy
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
		thisPolicy["bypass_cache_identifier"] = policy.BypassCacheIdentifier
		thisPolicy["content_type_filter"] = policy.ContentTypeFilter
		thisPolicy["header_filter"] = policy.HeaderFilter
		thisPolicy["method_filter"] = policy.MethodFilter
		thisPolicy["path_filter"] = policy.PathFilter
		thisPolicy["status_code_match"] = policy.StatusCodeMatch

		originPullPolicyIface = append(originPullPolicyIface, thisPolicy)
	}

	return originPullPolicyIface
}

// =========
// Request & Response Modifications

// OriginRequestModificationFromState ...
func (c *Configuration) OriginRequestModificationFromState(state []interface{}) error {
	modList := []*OriginRequestModification{}

	// extract
	for _, mod := range state {
		thisModRaw := mod.(map[string]interface{})
		thisMod := &OriginRequestModification{
			Enabled:     thisModRaw["enabled"].(bool),
			AddHeaders:  thisModRaw["add_headers"].(string),
			FlowControl: thisModRaw["flow_control"].(string),
		}
		modList = append(modList, thisMod)
	}

	if len(modList) > 0 {
		c.OriginRequestModification = modList
	}

	return nil
}

// OriginRequestModificationFromModel ...
func (c *Configuration) OriginRequestModificationFromModel() []interface{} {
	thisMap := make([]interface{}, 0)
	for _, mod := range c.OriginRequestModification {
		thisMap = append(thisMap, mod.Map())
	}
	if len(thisMap) < 1 {
		return nil
	}
	return thisMap
}

// OriginResponseModificationFromState ...
func (c *Configuration) OriginResponseModificationFromState(state []interface{}) error {
	modList := []*OriginResponseModification{}

	// extract
	for _, mod := range state {
		thisModRaw := mod.(map[string]interface{})
		thisMod := &OriginResponseModification{
			Enabled:     thisModRaw["enabled"].(bool),
			AddHeaders:  thisModRaw["add_headers"].(string),
			FlowControl: thisModRaw["flow_control"].(string),
		}
		modList = append(modList, thisMod)
	}

	if len(modList) > 0 {
		c.OriginResponseModification = modList
	}

	return nil
}

// OriginResponseModificationFromModel ...
func (c *Configuration) OriginResponseModificationFromModel() []interface{} {
	thisMap := make([]interface{}, 0)
	for _, mod := range c.OriginResponseModification {
		thisMap = append(thisMap, mod.Map())
	}
	if len(thisMap) < 1 {
		return nil
	}
	return thisMap
}

// ClientRequestModificationFromState ...
func (c *Configuration) ClientRequestModificationFromState(state []interface{}) error {
	modList := []*ClientRequestModification{}

	// extract
	for _, mod := range state {
		thisModRaw := mod.(map[string]interface{})
		thisMod := &ClientRequestModification{
			Enabled:     thisModRaw["enabled"].(bool),
			AddHeaders:  thisModRaw["add_headers"].(string),
			FlowControl: thisModRaw["flow_control"].(string),
		}
		modList = append(modList, thisMod)
	}

	if len(modList) > 0 {
		c.ClientRequestModification = modList
	}

	return nil
}

// ClientRequestModificationFromModel ...
func (c *Configuration) ClientRequestModificationFromModel() []interface{} {
	thisMap := make([]interface{}, 0)
	for _, mod := range c.ClientRequestModification {
		thisMap = append(thisMap, mod.Map())
	}
	if len(thisMap) < 1 {
		return nil
	}
	return thisMap
}

// ClientResponseModificationFromState ...
func (c *Configuration) ClientResponseModificationFromState(state []interface{}) error {
	modList := []*ClientResponseModification{}

	// extract
	for _, mod := range state {
		thisModRaw := mod.(map[string]interface{})
		thisMod := &ClientResponseModification{
			Enabled:     thisModRaw["enabled"].(bool),
			AddHeaders:  thisModRaw["add_headers"].(string),
			FlowControl: thisModRaw["flow_control"].(string),
		}
		modList = append(modList, thisMod)
	}

	if len(modList) > 0 {
		c.ClientResponseModification = modList
	}

	return nil
}

// ClientResponseModificationFromModel ...
func (c *Configuration) ClientResponseModificationFromModel() []interface{} {
	thisMap := make([]interface{}, 0)
	for _, mod := range c.ClientResponseModification {
		thisMap = append(thisMap, mod.Map())
	}

	return thisMap
}

// =========
// Delivery

// DeliveryFromModel ...
func (c *Configuration) DeliveryFromModel() {

}

// DeliveryFromState ...
func (c *Configuration) DeliveryFromState() {

}
