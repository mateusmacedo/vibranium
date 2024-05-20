package contract

// Value represents a generic value that can be of any type.
type Value any

// Validator is a generic interface that defines a validation method.
// Implementations of this interface should provide the logic to validate values of type T.
type Validator[T Value] interface {
	// Validate takes a value of type T and returns an error if the validation fails.
	// If the validation succeeds, it returns nil.
	//
	// Parameters:
	// - value: The value of type T to be validated.
	//
	// Returns:
	// - error: An error if the validation fails, or nil if the validation succeeds.
	Validate(value T) error
}

// ValidationFunc is a generic function type that implements the Validator interface.
// It defines a function that takes a value of type T and returns an error if the validation fails.
type ValidationFunc[T Value] func(value T) error

// Validate calls the ValidationFunc with the provided value and returns the result.
// This method allows ValidationFunc to satisfy the Validator interface.
//
// Parameters:
// - value: The value of type T to be validated.
//
// Returns:
// - error: An error if the validation fails, or nil if the validation succeeds.
func (f ValidationFunc[T]) Validate(value T) error {
	return f(value)
}
