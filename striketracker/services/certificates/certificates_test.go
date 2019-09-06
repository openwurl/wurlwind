package certificates

import (
	"context"
	"testing"
	"time"

	"github.com/openwurl/wurlwind/pkg/integration"
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
