package striketracker

import "fmt"

// TODO
// Implement optional user/pass access token authentication
// for non-permanent auth token usage

// Authorization manages the bearer and auth tokens
type Authorization struct {
	authorizationHeaderToken string
}

// GetBearer returns the bearer token header value
func (a *Authorization) GetBearer() string {
	return fmt.Sprintf("Bearer %s", a.authorizationHeaderToken)
}
