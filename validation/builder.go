package validation

import "github.com/mateusmacedo/vibranium/validation/contract"

// Builder is a generic builder that constructs a Composite validator for type T.
type Builder[T any] struct {
	composite *Composite[T] // composite is the composite validator being built.
}

// NewBuilder creates and returns a new Builder for type T.
// It initializes the composite validator.
func NewBuilder[T any]() *Builder[T] {
	return &Builder[T]{composite: NewComposite[T]()}
}

// Add adds a validator for a specific field to the composite validator.
// If the field already has validators, the new validator is appended to the list.
// It returns the builder to allow for method chaining.
//
// Parameters:
// - field: The name of the field to which the validator will be applied.
// - validator: The validator to be added.
//
// Returns:
// - *Builder[T]: The builder itself to allow for method chaining.
func (b *Builder[T]) Add(field string, validator contract.Validator[T]) *Builder[T] {
	b.composite.AddValidator(field, validator)
	return b
}

// Build finalizes and returns the composite validator being built.
//
// Returns:
// - *Composite[T]: The composite validator.
func (b *Builder[T]) Build() *Composite[T] {
	return b.composite
}
