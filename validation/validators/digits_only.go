package validators

import (
	"errors"
	"unicode"
)

// DigitsOnly is a validator that ensures a string contains only digit characters.
type DigitsOnly struct{}

// Validate checks if the given string contains only digit characters.
// If the string contains any non-digit characters, it returns an error.
// If the string contains only digits, it returns nil.
//
// Parameters:
// - value: The string to be validated.
//
// Returns:
// - error: An error if the string contains non-digit characters, or nil if the string is valid.
func (v DigitsOnly) Validate(value string) error {
	for _, char := range value {
		if !unicode.IsDigit(char) {
			return errors.New("value must contain only digits")
		}
	}
	return nil
}
