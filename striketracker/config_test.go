package striketracker

import (
	"testing"

	validator "gopkg.in/go-playground/validator.v9"
)

const (
	ConfigTestToken = "98fg798hu7sd9hu32kj23kj23"
	ConfigTestID    = "TestApplication"
)

func TestDefaults(t *testing.T) {

	config, err := NewConfiguration()

	if err != nil {
		t.Errorf("Did not expect error %v", err)
	}

	if config.Debug {
		t.Errorf("Expected Debug to be false")
	}

	if config.ApplicationID != "" {
		t.Errorf("Expected App Id to be empty, but was %v", config.ApplicationID)
	}

	if config.AuthorizationHeaderToken != "" {
		t.Errorf("Expected token to be empty, but was %v", config.AuthorizationHeaderToken)
	}
}

func TestSettingOptions(t *testing.T) {
	config, err := NewConfiguration(
		WithDebug(true),
		WithApplicationID(ConfigTestID),
		WithAuthorizationHeaderToken(ConfigTestToken),
	)

	if err != nil {
		t.Errorf("Did not expect error %v", err)
	}

	if !config.Debug {
		t.Errorf("Expected Debug to be true")
	}

	if config.ApplicationID != ConfigTestID {
		t.Errorf("Expected ApplicationId to be %v, but was %v", ConfigTestID, config.ApplicationID)
	}

	if config.AuthorizationHeaderToken != ConfigTestToken {
		t.Errorf("Expected AuthorizationHeaderToken to be %v, but was %v", ConfigTestToken, config.AuthorizationHeaderToken)
	}
}

func TestValidateDefaults(t *testing.T) {
	config, _ := NewConfiguration()
	err := config.Validate()

	if err == nil {
		t.Errorf("Expected configuration to not be valid, %v", err)
	}
}

func TestValidateMissingAuthToken(t *testing.T) {
	config, _ := NewConfiguration(WithApplicationID(ConfigTestID))
	err := config.Validate()

	if err == nil {
		t.Errorf("Expected config to not be valid")
	}

	// TODO extract the error to give more context
	if _, ok := err.(*validator.InvalidValidationError); ok {
		t.Errorf("Invalid Validation Error")
	}

	// TODO extract more context than just err.Field()
	for _, err := range err.(validator.ValidationErrors) {
		if err.Field() != "AuthorizationHeaderToken" {
			t.Errorf("Unexpected validation error on field %s", err.Field())
		}
	}
}
