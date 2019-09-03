package models

import "fmt"

// Response is the baseline response from the Striketracker API
type Response struct {
	Error string `json:"error,omitempty"`
	Code  int    `json:"code,omitempty"`
}

// Err returns the embedded error if it exists
func (r *Response) Err(err error) error {
	if r.Error != "" {
		embed := fmt.Sprintf("%d: %s", r.Code, r.Error)
		return fmt.Errorf("%s: (%s)", err.Error(), embed)
	}
	return err
}
