package models

import (
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
