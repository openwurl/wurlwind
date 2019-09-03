package striketracker

import (
	"testing"
)

const (
	ConfigTestToken = "98fg798hu7sd9hu32kj23kj23"
	ConfigTestID    = "TestApplication"
)

func TestDefaults(t *testing.T) {

	config, err := NewConfiguration()

	if err != nil {
		t.Errorf("Did not expect error")
	}

	if config.Debug {
		t.Errorf("Expected Debug to be false")
	}

	if config.ApplicationID != "" {
		t.Errorf("Expected App Id to be empty")
	}

	if config.AuthorizationHeaderToken != "" {
		t.Errorf("Expected token to be empty")
	}
}

func TestSettingOptions(t *testing.T) {
	config, err := NewConfiguration(
		WithDebug(true),
		WithApplicationID(ConfigTestID),
		WithAuthorizationHeaderToken(ConfigTestToken),
	)

	if err != nil {
		t.Errorf("Did not expect error")
	}

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

func TestValidateDefaults(t *testing.T) {
	config, _ := NewConfiguration()
	err := config.Validate()

	if err == nil {
		t.Errorf("Expected configuration to not be valid, %s", err)
	}
}
