package endpoints

import "fmt"

// Endpoint describes an API endpoint at striketracker
type Endpoint struct {
	BasePath BasePath
	Path     string
}

// String returns the base path up until account hash
func (e *Endpoint) String() string {
	return fmt.Sprintf("%s%s%s", URL, V1, e.BasePath)
}

// Format returns a formatted URL with account hash and path
func (e *Endpoint) Format(accountHash string) string {
	return fmt.Sprintf("%s/%s%s", e.String(), accountHash, e.Path)
}
