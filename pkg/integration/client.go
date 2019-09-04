package integration

import (
	"fmt"
	"os"

	"github.com/openwurl/wurlwind/striketracker"
)

// NewIntegrationClient returns a preconfigured integration client
// Requires AUTHORIZATIONHEADERKEY environment variable to be defined
// Optional APPLICATIONID environment variable to identify your application in logs
// To be used with integration tests only
func NewIntegrationClient() (*striketracker.Client, error) {
	authorizationHeaderToken := os.Getenv("AUTHORIZATIONHEADERKEY")
	if authorizationHeaderToken == "" {
		return nil, fmt.Errorf("No AUTHORIZATIONHEADERKEY defined, cannot run integration tests")
	}

	appID := os.Getenv("APPLICATIONID")
	if appID == "" {
		appID = "WurlWindIntegration"
	}

	// Configure the client
	c, err := striketracker.NewClientWithOptions(
		striketracker.WithApplicationID(appID),
		striketracker.WithDebug(true),
		striketracker.WithAuthorizationHeaderToken(authorizationHeaderToken),
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}
