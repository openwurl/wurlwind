package certificates

import (
	"context"
	"testing"
	"time"

	"github.com/openwurl/wurlwind/pkg/integration"
	"github.com/openwurl/wurlwind/striketracker/models"
)

// setup is called by tests to set up the client for integration testing
func setup() (*Service, string, error) {
	c, err := integration.NewIntegrationClient()
	if err != nil {
		return nil, "", err
	}

	s := New(c)

	accountHash, err := integration.GetIntegrationAccountHash()
	if err != nil {
		return nil, "", err
	}

	return s, accountHash, nil
}

// TestDestructiveCertificateSuiteIntegration tests create/update/get/delete
func TestDestructiveCertificateSuiteIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s, accountHash, err := setup()
	if err != nil {
		t.Fatalf("Expected client to be configured and account hash to be found but received error: %v", err)
	}

	cert, err := integration.GetCertificateIntegrationValues()
	if err != nil {
		t.Fatalf("Expected integration certificate data to be found but received error: %v", err)
	}

	var testSuite = []struct {
		name       string
		certCreate *models.Certificate
		certUpdate *models.Certificate
	}{
		{
			name: "Destructive Certificate Suite Test 01 without bundle and update with bundle",
			certCreate: &models.Certificate{
				Certificate: cert.Certificate,
				Key:         cert.Key,
			},
			certUpdate: &models.Certificate{
				CABundle:    cert.CABundle,
				Certificate: cert.Certificate,
				Key:         cert.Key,
			},
		},
	}

	// test with timeout context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 8*time.Second)
	defer cancel()

	// Run suite loop
	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			// Create
			t.Logf("Create %s", tt.name)

			createResponse, err := s.Upload(ctx, accountHash, tt.certCreate)
			if err != nil {
				t.Fatalf("error uploading certificate: %v", err)
			}
			if createResponse.ID == 0 {
				t.Fatalf("error uploading certificate, got no ID: %v", createResponse.ID)
			}

			/*
				These tests are all rather hard coded due to the nature of how we have
				to manage the certificate right now and how static certificates are

				This may cause problems in the future and a more graceful way of handling it
				should be investigated, since this test would likely fail if the cert expired
				or was not the exact cert defined for this test
			*/

			if createResponse.Issuer != "Let's Encrypt" {
				t.Fatalf("error, uploaded certificate issuer is Let's Encrypt but got: %s", createResponse.Issuer)
			}
			if createResponse.CertificateInformation.Subject.CN != "*.integrations.wurl.com" {
				t.Fatalf("error, expected Common Name *.integrations.wurl.com but got %s", createResponse.CertificateInformation.Subject.CN)
			}
			if createResponse.Trusted != false {
				t.Fatalf("error, expected certificate to NOT be trusted but it IS")
			}

			// Inject ID
			tt.certUpdate.ID = createResponse.ID

			// update
			t.Logf("Update %s with CA Bundle", tt.name)

			updateResponse, err := s.Update(ctx, accountHash, tt.certUpdate)
			if err != nil {
				t.Errorf("error updating certificate: %v", err)
			}

			if updateResponse.Trusted != true {
				t.Errorf("error, expected certificate to BE TRUSTED but it is NOT")
			}

			// Read
			t.Logf("Read %s", tt.name)
			readResponse, err := s.Get(ctx, accountHash, updateResponse.ID)
			if err != nil {
				t.Errorf("error fetching certificate: %v", err)
			}
			if readResponse.Fingerprint != updateResponse.Fingerprint {
				t.Errorf("Expected read and updated fingerprints to be identical, but was not")
			}
			if readResponse.CertificateInformation.Subject.CN != updateResponse.CertificateInformation.Subject.CN {
				t.Errorf("Expected read and updated CN to be identical, but was not")
			}

			// Delete
			t.Logf("Delete %s", tt.name)
			deleteErr := s.Delete(ctx, accountHash, readResponse.ID)
			if deleteErr != nil {
				t.Fatalf("error deleting certificate: %v", err)
			}
		})
	}

}

func TestListIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s, accountHash, err := setup()
	if err != nil {
		t.Fatalf("Expected client to be configured and account hash to be found but received error: %v", err)
	}

	// test with timeout context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	list, err := s.List(ctx, accountHash)
	if err != nil {
		t.Fatalf("Expected no error on List operation but got: %v", err)
	}

	t.Logf("Found %d certificates", len(list.List))
}
