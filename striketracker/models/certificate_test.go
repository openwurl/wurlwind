package models

import (
	"encoding/json"
	"testing"
)

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
