package striketracker

import (
	"fmt"
	"testing"
)

const (
	TestToken = "98fg798hu7sd9hu32kj23kj23"
	TestID    = "TestApplication"
)

var BaseConfiguration = &Configuration{
	ApplicationID:            TestID,
	AuthorizationHeaderToken: TestToken,
}

func TestNewClient(t *testing.T) {
	// Test client creation is rejected without secret
	conf := &Configuration{
		ApplicationID: TestID,
	}
	c, err := NewClient(conf)
	if err == nil {
		t.Errorf("expected an error for lacking auth secret but received none")
	}

	if c != nil {
		t.Error("expected returned client to be nil, but it managed to configure itself")
	}

	// Test client creation is rejected without ApplicationID
	conf = &Configuration{
		AuthorizationHeaderToken: TestToken,
	}
	c, err = NewClient(conf)
	if err == nil {
		t.Errorf("expected an error for lacking applicatoinID but received none")
	}

	if c != nil {
		t.Error("expected returned client to be nil, but it managed to configure itself")
	}

	// Test client configures itself successfully
	c, err = NewClient(BaseConfiguration)
	if err != nil {
		t.Error("Expected client to configure successfully but it failed")
	}

	if c.ApplicationID != TestID {
		t.Errorf("Expected ApplicationID to be %s but found %s", TestID, c.ApplicationID)
	}

	if c.Identity.AuthorizationHeaderToken != TestToken {
		t.Errorf("Expected AuthorizationHeaderToken to be %s but found %s", TestToken, c.Identity.AuthorizationHeaderToken)
	}

	expectedBearer := fmt.Sprintf("Bearer: %s", TestToken)
	foundBearer := c.Identity.GetBearer()
	if foundBearer != expectedBearer {
		t.Errorf("Expected client to produce bearer [%s] but got [%s]", expectedBearer, foundBearer)
	}

}
