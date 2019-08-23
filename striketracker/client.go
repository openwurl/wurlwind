package striketracker

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Header defines a header to be attached to requests
type Header struct {
	Key   string
	Value string
}

const ()

// Client is the primary wrapper for API interactions
type Client struct {
	Debug         bool
	Auth          *Authorization
	c             *http.Client
	ApplicationID string
	Headers       []*Header
}

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

	// Add auth token at this step to support refreshes
	req.Header.Set("Authorization", c.Auth.GetBearer())

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

// GetHeaders Generates the minimum requited headers
func (c *Client) GetHeaders() []*Header {
	var headers []*Header
	appID := &Header{
		Key:   "X-Application-Id",
		Value: c.ApplicationID,
	}
	headers = append(headers, appID)

	return headers
}
