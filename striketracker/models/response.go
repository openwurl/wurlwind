package models

import "fmt"

// Response is the baseline response from the Striketracker API
type Response struct {
	Message string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}

// Error returns the formatted error of the embedded if it exists
//
// You can use the provided errors to conveniently match ones you may expect
//
//	 if err != nil {
//	 	 if err == striketracker.ErrUnauthenticated {
//		 	 // handle
//		 } else if err == striketracker.ErrResourceExists {
//			 // handle
//		 }
//	 }
func (r *Response) Error() error {
	if r.Message != "" {
		return fmt.Errorf("%d: %s", r.Code, r.Message)
	}
	return nil
}
