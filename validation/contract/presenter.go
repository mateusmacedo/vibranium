package contract

import "github.com/mateusmacedo/vibranium/validation/errors"

// Presenter is an interface that defines a method for presenting validation errors.
// Implementations of this interface should define the logic to convert validation errors
// into an appropriate string format, such as JSON, XML, etc.
type Presenter interface {
	// Present receives a pointer to an Errors structure and returns a string.
	// The returned string represents the errors in an appropriate format.
	//
	// Parameters:
	// - errors: A pointer to an Errors structure containing the validation errors.
	//
	// Returns:
	// - string: A string representing the validation errors in an appropriate format.
	Present(errors *errors.Errors) string
}
