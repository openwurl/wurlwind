package validation

import validator "gopkg.in/go-playground/validator.v9"

// isDomain validates the field is a valid domain
func isDomain(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	return DomainRegExp.MatchString(val)
}

// validOriginPath validates that the origin path is acceptable
// Requires preceeding slash
func validOriginPath(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	// Not required
	if val == "" {
		return true
	}
	return OriginPathRexExp.MatchString(val)
}
