package striketracker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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
	c             *http.Client // TODO don't use DefaultClient and allow injectable
	ApplicationID string
	Headers       []*Header
}

// NewClientFromConfiguration will return a client with the provided configuration
func NewClientFromConfiguration(config *Configuration) (*Client, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	client, err := NewClient(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewClientWithOptions returns a configured client from functional parameters
func NewClientWithOptions(opts ...Option) (*Client, error) {
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

	if config.ApplicationID == "" {
		return nil, fmt.Errorf("ApplicationID is required - this is the only way to identify your requests in highwinds logs")
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

	// TODO eventually instantiate a custom client but not ready for that yet

	// Configure timeout on default client
	if config.Timeout == 0 {
		c.c.Timeout = time.Second * 10
	} else {
		c.c.Timeout = time.Second * time.Duration(config.Timeout)
	}

	// Set default headers
	c.Headers = c.GetHeaders()
	return c, nil
}

// NewRequestContext assembles a request with sensitive information and context
func (c *Client) NewRequestContext(ctx context.Context, method HTTPMethod, URL string, body interface{}) (*http.Request, error) {
	// Attach body if applicable
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
	req = req.WithContext(ctx)

	for _, header := range c.Headers {
		req.Header.Set(header.Key, header.Value)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
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

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
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
