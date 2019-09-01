package striketracker

import (
	"testing"
)

const (
	ConfigTestToken = "98fg798hu7sd9hu32kj23kj23"
	ConfigTestID    = "TestApplication"
)

func TestWithDebug(t *testing.T) {
	config := Configuration{}
	config.Option(WithDebug(true), WithApplicationID(ConfigTestID), WithAuthorizationHeaderToken(ConfigTestToken))
	if !config.Debug {
		t.Errorf("Expected Debug to be true")
	}

	if config.ApplicationID != ConfigTestID {
		t.Errorf("Expected ApplicationId to be %s", ConfigTestID)
	}

	if config.AuthorizationHeaderToken != ConfigTestToken {
		t.Errorf("Expected %s", ConfigTestToken)
	}
}
