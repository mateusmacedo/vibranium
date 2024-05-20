package validators

import "errors"

// StringNonEmpty is a validator that ensures a string is not empty.
type StringNonEmpty struct{}

// Validate checks if the given string is non-empty.
// If the string is empty, it returns an error.
// If the string is non-empty, it returns nil.
//
// Parameters:
// - value: The string to be validated.
//
// Returns:
// - error: An error if the string is empty, or nil if the string is non-empty.
func (v StringNonEmpty) Validate(value string) error {
	if value == "" {
		return errors.New("value cannot be empty")
	}
	return nil
}
