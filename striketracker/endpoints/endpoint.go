package endpoints

import "fmt"

// Endpoint describes a base API endpoint at striketracker
type Endpoint struct {
	BasePath BasePath
	Path     string
}

// String returns the base path up until account hash
func (e *Endpoint) String() string {
	return fmt.Sprintf("%s%s%s", URL, V1, e.BasePath)
}

// FormatAccountHash returns the base path up until the account hash
func (e *Endpoint) FormatAccountHash(accountHash string) string {
	return fmt.Sprintf("%s/%s", e.String(), accountHash)
}

// Format returns a formatted URL with account hash and path
func (e *Endpoint) Format(accountHash string) string {
	return fmt.Sprintf("%s/%s%s", e.String(), accountHash, e.Path)
}

// CustomFormat appends strings in order after account hash
func (e *Endpoint) CustomFormat(accountHash string, fields ...string) string {
	output := e.FormatAccountHash(accountHash)
	for field := range fields {
		tmp := fmt.Sprintf("%s/%s", output, fields[field])
		output = tmp
	}
	return output
}
