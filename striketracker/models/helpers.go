package models

import (
	"strconv"
)

// BoolFromInterface will attempt to extract a bool value from an interface in several ways
// and Will return fallback if no valid parsing can be achieved
func BoolFromInterface(data interface{}, fallback bool) bool {
	b, ok := data.(bool)
	if ok {
		return b
	}

	b, err := strconv.ParseBool(data.(string))
	if err != nil {
		return fallback
	}
	return b

}

// IntFromInterface will attempt to extract an int value from an interface in several ways
// and Will return fallback if no valid int parsing can be achieved
func IntFromInterface(data interface{}, fallback int) int {
	i, ok := data.(int)
	if ok {
		return i
	}

	i, err := strconv.Atoi(data.(string))
	if err != nil {
		return fallback
	}
	return i

}
