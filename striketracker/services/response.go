package services

import (
	"fmt"
	"net/http"
)

// ValidateResponse HTTP response
func ValidateResponse(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		return fmt.Errorf("%s", resp.Status)
	}
	return nil
}
