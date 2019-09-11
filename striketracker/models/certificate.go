package models

import (
	"fmt"
	"os"
	"reflect"

	"github.com/openwurl/wurlwind/pkg/fileio"
	"github.com/openwurl/wurlwind/pkg/validation"
	validator "gopkg.in/go-playground/validator.v9"
)

// CertificateResponse encapsulates a list of certificates
type CertificateResponse struct {
	List []Certificate `json:"list"`
}

// CertificateRequester represents the requesting entity of a certificate
type CertificateRequester struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"email"`
}

// Certificate encapsulates a TLS certificate request and response on a subaccount
//
// Certificate, CABundle, and the private Key can be all loaded in
// either from environment variables or env-var defined files
// via FromFile/FromEnv methods attached
//
type Certificate struct {
	Response
	CABundle               string                  `json:"caBundle"`                        // text of CA bundle
	Certificate            string                  `json:"certificate" validate:"required"` // text of x.509 cert
	CertificateInformation *CertificateInformation `json:"certificateInformation"`          // x.509 model
	Ciphers                string                  `json:"ciphers"`
	CommonName             string                  `json:"commonName"`
	CreatedDate            string                  `json:"createdDate"`
	ExpirationDate         string                  `json:"expirationDate"`
	Fingerprint            string                  `json:"fingerprint"`
	ID                     int                     `json:"id"`
	Issuer                 string                  `json:"issuer,omitempty"`
	Key                    string                  `json:"key" validate:"required"`
	Requester              *CertificateRequester   `json:"certificateRequester"`
	Trusted                bool                    `json:"trusted"`
	UpdatedDate            string                  `json:"updatedDate"`
}

// CABundleFromFile attaches a CA bundle from the given file
func (c *Certificate) CABundleFromFile(filepath string) error {
	contents, err := fileio.FileToString(filepath)
	if err != nil {
		return err
	}
	c.CABundle = contents
	return nil
}

// CABundleFromEnv attaches a CA bundle from the environment variable
func (c *Certificate) CABundleFromEnv(evar string) error {
	caBundle := os.Getenv(evar)
	if caBundle == "" {
		return fmt.Errorf("Env var %s does not contain data", evar)
	}

	c.CABundle = caBundle
	return nil
}

// CertificateFromFile attaches the certificate from the given file
func (c *Certificate) CertificateFromFile(filepath string) error {
	contents, err := fileio.FileToString(filepath)
	if err != nil {
		return err
	}
	c.Certificate = contents
	return nil

}

// CertificateFromEnv attaches a Certificate from the environment variable
func (c *Certificate) CertificateFromEnv(evar string) error {
	cert := os.Getenv(evar)
	if cert == "" {
		return fmt.Errorf("Env var %s does not contain data", evar)
	}

	c.Certificate = cert
	return nil
}

// KeyFromFile attaches the key from the given file
func (c *Certificate) KeyFromFile(filepath string) error {
	contents, err := fileio.FileToString(filepath)
	if err != nil {
		return err
	}
	c.Key = contents
	return nil
}

// KeyFromEnv attaches a Private Key from the environment variable
func (c *Certificate) KeyFromEnv(evar string) error {
	key := os.Getenv(evar)
	if key == "" {
		return fmt.Errorf("Env var %s does not contain data", evar)
	}

	c.Key = key
	return nil
}

// Validate validates the Certificate struct data
func (c *Certificate) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(c); err != nil {
		return err
	}

	return nil
}

// CertificateInformation encapsulates a debundled cert
type CertificateInformation struct {
	Name    string              `json:"name"`
	Subject *CertificateSubject `json:"subject"`
}

// CertificateSubject is a sub field within CertificateInformation
type CertificateSubject struct {
	CN string `json:"CN"`
}

// CertificateHostsResponse is a list of hosts consuming a certificate in its native format
// from the API
type CertificateHostsResponse map[string][]*CertificateHost

// Process returns a CertificateHosts from a CertificateHostsResponse
func (c *CertificateHostsResponse) Process() (*CertificateHosts, error) {
	keys := reflect.ValueOf(*c).MapKeys()
	if len(keys) < 1 || len(keys) > 1 {
		return nil, fmt.Errorf("There should only be a single certificate")
	}
	chs := &CertificateHosts{
		CertificateName: keys[0].String(),
	}
	for _, host := range (*c)[chs.CertificateName] {
		chs.Hosts = append(chs.Hosts, host)
	}
	return chs, nil
}

// CertificateHosts is a list of hosts consuming a certificate in structured format
type CertificateHosts struct {
	Response
	CertificateName string
	Hosts           []*CertificateHost
}

// CertificateHost is a host consuming a certificate
type CertificateHost struct {
	ID          int    `json:"id,string"`
	Hostname    string `json:"hostname"`
	HashCode    string `json:"hashCode"`
	VIP         string `json:"vip"`
	AccountHash string `json:"accountHash"`
	VIPName     string `json:"vipName"`
	Services    string `json:"services"`
}
