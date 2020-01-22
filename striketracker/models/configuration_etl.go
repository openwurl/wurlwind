package models

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
