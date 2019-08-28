// Package validate provides a central
// way to validate individual fields and different
// implementations of config
package validate

import (
	"regexp"
	"strconv"
	"strings"

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
	err := cv.validator.RegisterValidation("hlsdate", isWurlDate)
	if err != nil {
		panic(err)
	}
	err = cv.validator.RegisterValidation("apiversion", isWurlAPIVersion)
	if err != nil {
		panic(err)
	}
	err = cv.validator.RegisterValidation("domain", isDomain)
	if err != nil {
		panic(err)
	}
}

// Validate is the main validation step for structs
func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// isWurlDate determines if a date conforms to
// YYYY-MM-DD format
func isWurlDate(fl validator.FieldLevel) bool {
	// expect xxxx-xx-xx
	val := fl.Field().String()
	split := strings.Split(val, "-")
	if len(split) < 3 || len(split) > 3 {
		return false
	}

	for i, part := range split {
		cvt, err := strconv.Atoi(part)
		if err != nil {
			return false
		}

		switch i {
		case 1:
			if cvt < 0 || cvt > 12 {
				return false
			}
		case 2:
			if cvt < 0 || cvt > 31 {
				return false
			}
		}

	}

	return true
}

// isWurlAPIVersion determines if the field conforms
// to VV.v format
func isWurlAPIVersion(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	split := strings.Split(val, ".")
	if len(split) < 2 || len(split) > 2 {
		return false
	}

	for _, part := range split {
		_, err := strconv.Atoi(part)
		if err != nil {
			return false
		}
	}

	return true
}

// isDomain validates the field is a valid domain
func isDomain(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	return DomainRegExp.MatchString(val)
}
