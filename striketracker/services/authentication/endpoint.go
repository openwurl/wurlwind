package authentication

import (
	"fmt"

	"github.com/openwurl/wurlwind/striketracker/endpoints"
)

// Endpoint storage since auth varies heavily
const (
	UserEndpoint         = "%s/users/%s/tokens"
	TokenEndpoint        = "%s/users/%s/tokens/%s"
	AuthenticateEndpoint = "/auth/token"
)

// AuthEndpoint defines the custom endpoint for authentication
type AuthEndpoint struct {
	*endpoints.Endpoint
}

func (a *AuthEndpoint) formatUser(accountHash string, userID string) string {
	return fmt.Sprintf(UserEndpoint, a.FormatAccountHash(accountHash), userID)
}

func (a *AuthEndpoint) formatToken(accountHash string, userID string, token string) string {
	return fmt.Sprintf(TokenEndpoint, a.FormatAccountHash(accountHash), userID, token)
}

func (a *AuthEndpoint) formatAuth() string {
	// This is a unique endpoint not following structure
	return fmt.Sprintf("%s%s", endpoints.URL, AuthenticateEndpoint)
}
