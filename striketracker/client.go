package striketracker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/openwurl/wurlwind/striketracker/identity"
)

// Header defines a header to be attached to requests
type Header struct {
	Key   string
	Value string
}

// Client is the primary wrapper for API interactions
type Client struct {
	Debug bool
	//Auth          *auth.Wrapper
	Identity      *identity.Identification
	c             *http.Client
	ApplicationID string
	Headers       []*Header
}

// NewClientWithOptions returns a configured client from functional parameters
func NewClientWithOptions(opts ...Config) (*Client, error) {
	options := &Configuration{}
	for _, opt := range opts {
		opt(options)
	}

	err := options.Validate()
	if err != nil {
		return nil, err
	}

	c, err := NewClient(options)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// NewClient returns a configured client
func NewClient(config *Configuration) (*Client, error) {
	// Check that authorization values are defined at all
	if config.AuthorizationHeaderToken == "" {
		return nil, fmt.Errorf("No authorization is defined. You need AuthorizationHeaderToken")
	}

	// Configure the client from final configuration
	c := &Client{
		c:             http.DefaultClient,
		Debug:         config.Debug,
		ApplicationID: config.ApplicationID,
		Identity: &identity.Identification{
			AuthorizationHeaderToken: config.AuthorizationHeaderToken,
		},
	}
	// Set default headers
	c.Headers = c.GetHeaders()
	return c, nil
}

/*
// NewClient returns a configured client
func NewClient(config *Config) *Client {
	c := &Client{
		c:     http.DefaultClient,
		Debug: config.Debug,
		//		Auth: &Authorization{
		//			authorizationHeaderToken: config.AuthorizationHeaderToken,
		//		},
		ApplicationID: config.ApplicationID,
	}
	c.Headers = c.GetHeaders()
	return c
}

striketracker.NewClient(
	WithConfig(&striketracker.Config{Debug: false}),
	WithEnv(),

)

// NewClient returns a configured client
func NewClient(debug bool, authorizationHeaderToken string, applicationID string) *Client {
	c := &Client{
		c:     http.DefaultClient,
		Debug: debug,
		Auth: &Authorization{
			authorizationHeaderToken: authorizationHeaderToken,
		},
		ApplicationID: applicationID,
	}

	c.Headers = c.GetHeaders()

	return c
}

// NewClientFromConfiguration returns a configured client from the given configuration
func NewClientFromConfiguration(config *Config) *Client {
	c := &Client{
		c:     http.DefaultClient,
		Debug: config.Debug,
		Auth: &Authorization{
			authorizationHeaderToken: config.AuthorizationHeaderToken,
		},
		ApplicationID: config.ApplicationID,
	}
	c.Headers = c.GetHeaders()
	return c
}

// NewClientFromEnv returns a configured client from env vars
func NewClientFromEnv() *Client {
	return nil
}
*/

// CreateRequest assembles a request with sensitive information
func (c *Client) CreateRequest(method HTTPMethod, URL string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method.String(), URL, buf)
	if err != nil {
		return nil, err
	}

	for _, header := range c.Headers {
		req.Header.Set(header.Key, header.Value)
	}

	// Add auth token from memory if it exists
	if c.Identity.AuthorizationHeaderToken != "" {
		req.Header.Set("Authorization", c.Identity.GetBearer())
	}

	return req, nil
}

// DoRequest performs the request
func (c *Client) DoRequest(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

// GetHeaders Generates the minimum required headers
func (c *Client) GetHeaders() []*Header {
	var headers []*Header
	appID := &Header{
		Key:   "X-Application-Id",
		Value: c.ApplicationID,
	}
	headers = append(headers, appID)

	return headers
}
