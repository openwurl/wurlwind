// Package integration provides helpers for integration tests
package integration

import (
	"fmt"
	"os"
)

// GetIntegrationAccountHash fetches the integration account hash from Env var as your
// target account to test against
func GetIntegrationAccountHash() (string, error) {
	accountHash := os.Getenv("INTEGRATIONACCOUNTHASH")
	if accountHash == "" {
		return "", fmt.Errorf("INTEGRATIONACCOUNTHASH must be defined to run integration tests")
	}

	return accountHash, nil
}
