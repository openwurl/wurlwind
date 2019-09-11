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

	if cert.Key == "" {
		return nil, fmt.Errorf("Unable to configure certificate private key from env vars or file")
	}
	if cert.Certificate == "" {
		return nil, fmt.Errorf("Unable to configure certificate from env vars or file")
	}
	if cert.CABundle == "" {
		return nil, fmt.Errorf("Unable to configure certificate CA Bundle from env vars or file")
	}

	return cert, nil
}

// GetCertificateIntegrationValues fetches the integration certificate from Env var
//func GetCertificateIntegrationValues() (string, string, string, error) {
/*
	privKey, err := getEnvVar("INTEGRATIONPRIVATEKEY")
	if err != nil {
		return
	}
*/
/*
	privKeyFile := os.Getenv("INTEGRATIONPRIVATEKEY")
	if privKeyFile == "" {
		return "", "", "", fmt.Errorf("INTEGRATIONPRIVATEKEY must be defined to run integration tests")
	}

	certFile := os.Getenv("INTEGRATIONCERTIFICATE")
	if certFile == "" {
		return "", "", "", fmt.Errorf("INTEGRATIONCERTIFICATE must be defined to run integration tests")
	}

	bundleFile := os.Getenv("INTEGRATIONBUNDLE")
	if bundleFile == "" {
		return "", "", "", fmt.Errorf("INTEGRATIONBUNDLE must be defined to run integration tests")
	}

	privKey, err := fileio.FileToString(privKeyFile)
	if err != nil {
		return "", "", "", err
	}
	cert, err := fileio.FileToString(certFile)
	if err != nil {
		return "", "", "", err
	}
	bundle, err := fileio.FileToString(bundleFile)
	if err != nil {
		return "", "", "", err
	}

	return privKey, cert, bundle, nil
*/
//}

func getEnvVar(evar string) (string, error) {
	this := os.Getenv(evar)
	if this == "" {
		return "", fmt.Errorf("%s must be defined", evar)
	}
	return this, nil
}
