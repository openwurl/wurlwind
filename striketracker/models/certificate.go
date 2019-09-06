package models

import (
	"fmt"
	"reflect"

	"github.com/openwurl/wurlwind/pkg/validation"
	validator "gopkg.in/go-playground/validator.v9"
)

// CertificateResponse encapsulates a list of certificates
type CertificateResponse struct {
	List []Certificate `json:"list"`
}

// Certificate encapsulates a TLS certificate on a subaccount
type Certificate struct {
	CABundle               string                  `json:"caBundle"`    // text of CA bundle
	Certificate            string                  `json:"certificate"` // text of x.509 cert
	CertificateInformation *CertificateInformation `json:"certificateInformation"`
	Ciphers                string                  `json:"ciphers"`
	CommonName             string                  `json:"commonName"`
	CreatedDate            string                  `json:"createdDate"`
	ExpirationDate         string                  `json:"expirationDate"`
	Fingerprint            string                  `json:"fingerprint"`
	ID                     int                     `json:"id"`
	Issuer                 string                  `json:"issuer"`
	Key                    string                  `json:"key"`
	Requester              string
	Trusted                bool   `json:"trusted"`
	UpdatedDate            string `json:"updatedDate"`
}

// Validate validates the struct data
func (c *Certificate) Validate() error {
	v := validation.NewValidator(validator.New())
	if err := v.Validate(c); err != nil {
		return err
	}

	return nil
}

// CertificateInformation encapsulates a debundled cert
type CertificateInformation struct {
	Certificate string `json:"certificate"`
	Key         string `json:"key"`
	CABundle    string `json:"caBundle"`
}

// CertificateHostsUnload is a list of hosts consuming a certificate in its native format
type CertificateHostsUnload map[string][]*CertificateHost

// Process returns a CertificateHosts
func (c *CertificateHostsUnload) Process() (*CertificateHosts, error) {
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
