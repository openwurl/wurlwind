package models

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/openwurl/wurlwind/pkg/debug"
	"github.com/openwurl/wurlwind/striketracker/models"
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
// Origin Pull Cache Extension

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
// Origin Pull Policy

// OriginPullPolicyFromState ...
func (c *Configuration) OriginPullPolicyFromState(state []interface{}) error {
	orderedList := make([]interface{}, len(state))
	c.OriginPullPolicy = &models.OriginPullPolicy{}

	// order the list by defined weight
	for _, policy := range state {
		policyCast := policy.(map[string]interface{})
		policyIndex := policyCast["weight"].(int)
		// TODO: Rest of keys

		if orderedList[policyIndex] != nil {
			return fmt.Errorf("Weight %d used multiple times", policyIndex)
		}
		orderedList[policyIndex] = policyCast
	}

	for _, policy := range orderedList {
		thisPolicy := models.NewOriginPullPolicyFromState(policy)
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
		// TODO: Rest of keys

		originPullPolicyIface = append(originPullPolicyIface, thisPolicy)
	}

	return originPullPolicyIface
}
