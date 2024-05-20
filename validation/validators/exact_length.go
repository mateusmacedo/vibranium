package validators

import (
	"errors"
	"fmt"
)

// ExactLength is a validator that ensures a string is exactly a specified length.
type ExactLength struct {
	Length int // Length is the exact number of characters the string must have.
}

// Validate checks if the given string is exactly the specified length.
// If the string length does not match, it returns an error.
// If the string length matches, it returns nil.
//
// Parameters:
// - value: The string to be validated.
//
// Returns:
// - error: An error if the string length does not match, or nil if the string length is valid.
func (v ExactLength) Validate(value string) error {
	if len(value) != v.Length {
		return errors.New(fmt.Sprintf("value must be exactly %d characters long", v.Length))
	}
	return nil
}
