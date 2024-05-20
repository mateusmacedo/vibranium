package validators

import "errors"

// PositiveNumber is a validator that ensures a number is positive.
type PositiveNumber struct{}

// Validate checks if the given integer is positive.
// If the integer is not positive (i.e., zero or negative), it returns an error.
// If the integer is positive, it returns nil.
//
// Parameters:
// - value: The integer to be validated.
//
// Returns:
// - error: An error if the integer is not positive, or nil if the integer is valid.
func (v PositiveNumber) Validate(value int) error {
	if value <= 0 {
		return errors.New("value must be positive")
	}
	return nil
}
