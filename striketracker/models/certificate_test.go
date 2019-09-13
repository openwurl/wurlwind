package models

import (
	"encoding/json"
	"os"
	"testing"
)

var (
	testFakeBundle = "FAKE BUNDLE"
	testFakeCert   = "FAKE CERT"
	testFakeKey    = "FAKE KEY"
)

func TestCertificate(t *testing.T) {
	os.Setenv("TESTKEY", testFakeKey)
	os.Setenv("TESTCERT", testFakeCert)
	os.Setenv("TESTBUNDLE", testFakeBundle)

	c := &Certificate{}
	err := c.KeyFromEnv("TESTKEY")
	if err != nil {
		t.Fatalf("Expected to fetch KEY environment variable but got error: %v", err)
	}
	err = c.CABundleFromEnv("TESTBUNDLE")
	if err != nil {
		t.Fatalf("Expected to fetch BUNDLE environment variable but got error: %v", err)
	}
	err = c.CertificateFromEnv("TESTCERT")
	if err != nil {
		t.Fatalf("Expected to fetch CERT environment variable but got error: %v", err)
	}

	if c.Key != testFakeKey {
		t.Fatalf("expected key %s but got %s", testFakeKey, c.Key)
	}
	if c.Certificate != testFakeCert {
		t.Fatalf("expected Cert %s but got %s", testFakeCert, c.Certificate)
	}
	if c.CABundle != testFakeBundle {
		t.Fatalf("expected CABundle %s but got %s", testFakeBundle, c.CABundle)
	}

	if err = c.Validate(); err != nil {
		t.Fatalf("expected certificate to pass validation for required fields but did not: %v", err)
	}

}

// TODO make this test more structured, hard to do with the byte payload input
func TestCertificateHost(t *testing.T) {
	test := []byte(`{
		"*.somedomain.com": [
			{
				"id": "8675309",
				"hostname": "wurl.somedns.net",
				"hashCode": "f87923j",
				"vip": null,
				"accountHash": "fd8923j",
				"vipName": null,
				"services": "4,40"
			}
		]
	}`)
	//var dest CertificateHosts
	var dest CertificateHostsResponse

	err := json.Unmarshal(test, &dest)
	if err != nil {
		t.Fatalf("Expected no error on unmarshal but received: %v", err)
	}

	if len(dest) > 1 {
		t.Fatalf("Expected no more than one certificate to match, got %d", len(dest))
	}

	//var certHost *CertificateHosts
	chs, err := dest.Process()
	if err != nil {
		t.Fatalf("Expected certificate hosts to process but received error: %v", err)
	}

	if chs.CertificateName != "*.somedomain.com" {
		t.Fatalf("Expected CertificateName *.somedomain.com but got %s", chs.CertificateName)
	}

	if chs.Hosts[0].ID != 8675309 {
		t.Fatalf("Expected host ID to be 8675309 but got %d", chs.Hosts[0].ID)
	}
}
