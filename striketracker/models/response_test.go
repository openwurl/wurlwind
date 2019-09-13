package models

import (
	"testing"

	"github.com/openwurl/wurlwind/striketracker"
)

func TestResponseErr(t *testing.T) {
	r := &Response{
		Code:    303,
		Message: "The authenticated user does not have an associated account",
	}

	err := r.Error()
	if err.Error() != striketracker.ErrUnassociatedUser {
		t.Fatalf("expected error %s but got %v", striketracker.ErrUnassociatedUser, err)
	}
}
