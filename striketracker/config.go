package striketracker

import (
	"github.com/wurlinc/hls-config/pkg/validate"
	validator "gopkg.in/go-playground/validator.v9"
)

// Configuration provides a service configuration for the client
type Configuration struct {
	Debug                    bool   `json:"debug"`
	AuthorizationHeaderToken string `json:"authorizationHeaderToken" validate:"required"`
	ApplicationID            string `json:"applicationID" validate:"required"`
}

// Validate validates the configuration is valid
func (c *Configuration) Validate() error {
	v := validate.NewValidator(validator.New())
	if err := v.Validate(c); err != nil {
		return err
	}
	return nil
}

// Config is a functional API for configuring the client
type Config func(*Configuration)

// WithDebug adds debug on instantiation
func WithDebug(debug bool) Config {
	return func(c *Configuration) {
		c.Debug = debug
	}
}

// WithAuthorizationHeaderToken adds an auth token on instantiation
func WithAuthorizationHeaderToken(token string) Config {
	return func(c *Configuration) {
		c.AuthorizationHeaderToken = token
	}
}

// WithApplicationID adds the ApplicationID on instantiation
func WithApplicationID(appID string) Config {
	return func(c *Configuration) {
		c.ApplicationID = appID
	}
}

/* Not Implemented yet
// WithConfigFile loads configuration from a configuration file
func WithConfigFile(filepath string) Config {
	return func(c *Configuration) {
	}
}

*/
