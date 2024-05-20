package validation

import (
	"github.com/mateusmacedo/vibranium/validation/contract"
	"github.com/mateusmacedo/vibranium/validation/errors"
)

// Composite is a generic type that holds multiple validators for type T.
type Composite[T any] struct {
	validators []WithField[T] // validators is a slice of WithField containing field names and their respective validators.
}

// WithField associates a field name with a validator for type T.
type WithField[T any] struct {
	Field     string                // Field is the name of the field to be validated.
	Validator contract.Validator[T] // Validator is the validator to be applied to the field.
}

// NewComposite creates and returns a new Composite for type T.
// It initializes the validators slice.
//
// Returns:
// - *Composite[T]: A pointer to the newly created Composite.
func NewComposite[T any]() *Composite[T] {
	return &Composite[T]{validators: []WithField[T]{}}
}

// AddValidator adds a validator for a specific field to the Composite.
// It takes the field name and the validator to be added as arguments.
//
// Parameters:
// - field: The name of the field to which the validator will be applied.
// - validator: The validator to be added.
func (c *Composite[T]) AddValidator(field string, validator contract.Validator[T]) {
	c.validators = append(c.validators, WithField[T]{Field: field, Validator: validator})
}

// Validate checks if the given value is valid using all the validators in the Composite.
// It returns an error if any validator fails. The errors are aggregated into an Errors structure.
//
// Parameters:
// - value: The value to be validated.
//
// Returns:
// - error: An error if any validator fails, or nil if all validators pass.
func (c *Composite[T]) Validate(value T) error {
	errs := &errors.Errors{}
	for _, vc := range c.validators {
		if err := vc.Validator.Validate(value); err != nil {
			if nestedErrs, ok := err.(*errors.Errors); ok {
				errs.AddNested(vc.Field, nestedErrs)
			} else {
				errs.Add(vc.Field, err.Error())
			}
		}
	}
	if errs.IsEmpty() {
		return nil
	}
	return errs
}
