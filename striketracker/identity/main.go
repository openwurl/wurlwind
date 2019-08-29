// Package identity defines configuration around
// authentication for the user interfacing with the API
package identity

import "fmt"

// Identification defines the user
type Identification struct {
	AuthorizationHeaderToken string
}

// SetToken sets the header token
func (i *Identification) SetToken(token string) {
	i.AuthorizationHeaderToken = token
}

// GetBearer returns the formatted bearer token
func (i *Identification) GetBearer() string {
	return fmt.Sprintf("Bearer %s", i.AuthorizationHeaderToken)
}
