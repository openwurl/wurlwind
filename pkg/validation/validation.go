package validation

import (
	"regexp"

	validator "gopkg.in/go-playground/validator.v9"
)

// RegExp compilations
var (
	DomainRegExp = regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
		]{2,3})$`)
)

// Validator is a custom validator
type Validator struct {
	validator *validator.Validate
}

// NewDefaultValidator ...
func NewDefaultValidator() *Validator {
	return NewValidator(validator.New())
}

// NewValidator constructs a custom validator with all custom validators registered.
func NewValidator(v *validator.Validate) *Validator {
	cv := &Validator{
		validator: v,
	}
	cv.RegisterCustom()
	return cv
}

// RegisterCustom registers all custom validators for use.
func (cv *Validator) RegisterCustom() {
	err := cv.validator.RegisterValidation("domain", isDomain)
	if err != nil {
		panic(err)
	}
}

// Validate is the main entry to validate structs
func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
