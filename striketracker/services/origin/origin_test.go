package origin

import (
	"testing"

	"github.com/openwurl/wurlwind/pkg/integration"
)

var (
	integrationAccountHash = "z3d5t6j7"
)

func TestGetIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	c, err := integration.NewIntegrationClient()
	if err != nil {
		t.Fatalf("Expected client to be configured but received error: %v", err)
	}

	s := New(c)

	var tests = []struct {
		name          string
		originID      string
		shouldSucceed bool
	}{
		{"Test should fail", "123456", false},
		{"Test get [Integration Test Origin For Get]", "235473", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thisOrigin, err := s.Get(integrationAccountHash, tt.originID)
			if tt.shouldSucceed {
				if err != nil {
					t.Fatalf("Expected to find origin with ID %s, but got error: %v", tt.originID, err)
				}
			} else {
				if err == nil {
					t.Fatalf("Expected to not find origin with ID %s, but received: %v", tt.originID, thisOrigin)
				}
			}
		})
	}
}
