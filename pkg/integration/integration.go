// Package integration provides helpers for integration tests
package integration

import (
	"fmt"
	"os"

	"github.com/openwurl/wurlwind/striketracker/models"
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

// GetCertificateIntegrationValues fetches the integration certificate from Env var
func GetCertificateIntegrationValues() (*models.Certificate, error) {
	var (
		envPrivKey     = "INTEGRATIONPRIVATEKEY"
		envPrivKeyFile = "INTEGRATIONPRIVATEKEYFILE"
		envCert        = "INTEGRATIONCERT"
		envCertFile    = "INTEGRATIONCERTFILE"
		envBundle      = "INTEGRATIONBUNDLE"
		envBundleFile  = "INTEGRATIONBUNDLEFILE"
	)

	cert := &models.Certificate{}

	// Either load in from file or env var
	if privKeyFile := os.Getenv(envPrivKeyFile); privKeyFile != "" {
		err := cert.KeyFromFile(privKeyFile)
		if err != nil {
			return nil, err
		}
	} else {
		err := cert.KeyFromEnv(envPrivKey)
		if err != nil {
			return nil, err
		}
	}

	if certFile := os.Getenv(envCertFile); certFile != "" {
		err := cert.CertificateFromFile(certFile)
		if err != nil {
			return nil, err
		}
	} else {
		err := cert.CertificateFromEnv(envCert)
		if err != nil {
			return nil, err
		}
	}

	if bundleFile := os.Getenv(envBundleFile); bundleFile != "" {
		err := cert.CABundleFromFile(bundleFile)
		if err != nil {
			return nil, err
		}
	} else {
		err := cert.CABundleFromEnv(envBundle)
		if err != nil {
			return nil, err
		}
	}

	if err := cert.Validate(); err != nil {
		return nil, err
	}

	return cert, nil
}
