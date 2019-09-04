package integration

import (
	"fmt"
	"os"

	"github.com/openwurl/wurlwind/striketracker"
)

// NewIntegrationClient returns a preconfigured integration client
// Requires AUTHORIZATIONHEADERKEY environment variable to be defined
// To be used with integration tests only
func NewIntegrationClient() (*striketracker.Client, error) {
	authorizationHeaderToken := os.Getenv("AUTHORIZATIONHEADERKEY")
	if authorizationHeaderToken == "" {
		return nil, fmt.Errorf("No AUTHORIZATIONHEADERKEY defined, cannot run integration tests")
	}

	// Configure the client
	c, err := striketracker.NewClientWithOptions(
		striketracker.WithApplicationID("WurlWindIntegration"),
		striketracker.WithDebug(true),
		striketracker.WithAuthorizationHeaderToken(authorizationHeaderToken),
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}
