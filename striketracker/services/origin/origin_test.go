package origin

import (
	"context"
	"reflect"
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

// TestDestructiveOriginSuiteIntegration tests create/update/get/delete
func TestDestructiveOriginSuiteIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	var testSuite = []struct {
		name   string
		create *models.Origin
		update *models.Origin
	}{
		{
			name: "Destructive Suite Test 01 with path update",
			create: &models.Origin{
				Name:     "CUGD Integration Test 01",
				Hostname: "integration.noroute.com",
				Port:     80,
			},
			update: &models.Origin{
				Name:     "CUGD Integration Test 01",
				Hostname: "integration.noroute.com",
				Port:     80,
				Path:     "/updated/with/path",
			},
		},
		{
			name: "Destructive Suite Test 02 with path and port update",
			create: &models.Origin{
				Name:     "CUGD Integration Test 02",
				Hostname: "integration.nohost.com",
				Port:     80,
				Path:     "/created/with/path",
			},
			update: &models.Origin{
				Name:     "CUGD Integration Test 02",
				Hostname: "differenthost.nohost.com",
				Port:     8080,
				Path:     "",
			},
		},
	}

	// Run setup
	s, accountHash, err := setup()
	if err != nil {
		t.Fatalf("Expected client to be configured and account hash to be found but received error: %v", err)
	}

	// Run suite in a loop
	// Create, Update, Get, Delete
	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			// Create
			t.Logf("Create %s", tt.name)

			// test with timeout context
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()

			createResponse, err := s.Create(ctx, accountHash, tt.create)
			if err != nil {
				t.Fatalf("Expected creation of integration origin but received error: %v", err)
			}

			if createResponse == nil {
				t.Fatalf("Expected to receive struct containing created origin details, but got nil")
			}

			if createResponse.ID == 0 {
				t.Fatalf("Did not receive an Origin ID from create response")
			}

			if createResponse.Name != tt.create.Name {
				t.Fatalf("Expected created origin to have name %s but has %s", tt.create.Name, createResponse.Name)
			}

			if createResponse.Hostname != tt.create.Hostname {
				t.Fatalf("Expected created origin to point to hostname %s but has %s", tt.create.Hostname, createResponse.Hostname)
			}

			if createResponse.Port != tt.create.Port {
				t.Fatalf("Expected created origin to use port %d but got %d", tt.create.Port, createResponse.Port)
			}

			if tt.create.Path == "" {
				if createResponse.Path != "/" {
					t.Fatalf("Expected no path on created origin but found %s", createResponse.Path)
				}
			} else {
				if createResponse.Path != tt.create.Path {
					t.Fatalf("Expected path on created origin to be %s but found %s", tt.create.Path, createResponse.Path)
				}
			}

			// Update
			t.Logf("Update %s", tt.name)
			update := *tt.update
			update.ID = createResponse.ID
			updatedResponse, err := s.Update(ctx, accountHash, &update)
			if err != nil {
				t.Fatalf("Expected update of integration origin but received error: %v", err)
			}

			if updatedResponse == nil {
				t.Fatalf("Expected to receive struct containing updated origin details, but got nil")
			}

			if tt.update.Path == "" {
				if updatedResponse.Path != "/" {
					t.Fatalf("Expected no path on updated origin but found %s", createResponse.Path)
				}
			} else {
				if updatedResponse.Path != tt.update.Path {
					t.Fatalf("Expected updated path to be [%s] but got [%s]", tt.update.Path, updatedResponse.Path)
				}
			}

			if updatedResponse.Port != tt.update.Port {
				t.Fatalf("Expected updated port to be %d but got %d", tt.update.Port, updatedResponse.Port)
			}

			// Get
			t.Logf("Get %s", tt.name)
			receivedResponse, err := s.Get(ctx, accountHash, updatedResponse.ID)
			if err != nil {
				t.Fatalf("Expected integration origin but received error: %v", err)
			}

			if !reflect.DeepEqual(receivedResponse, updatedResponse) {
				t.Fatalf("Expected update and get resposes to be identical, but they were not")
			}

			// Delete
			t.Logf("Delete %s", tt.name)
			err = s.Delete(ctx, accountHash, updatedResponse.ID)
			if err != nil {
				t.Fatalf("Expected deletion of integration origin but received error: %v", err)
			}
		})
	}

}
