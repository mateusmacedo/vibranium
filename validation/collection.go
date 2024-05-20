package validation

import (
	"fmt"

	"github.com/mateusmacedo/vibranium/validation/contract"
	"github.com/mateusmacedo/vibranium/validation/errors"
)

// Collection is a generic type that validates a collection of items of type T.
type Collection[T any] struct {
	itemValidator contract.Validator[T] // itemValidator is the validator used to validate each item in the collection.
}

// NewCollection creates and returns a new Collection with the specified item validator.
// The item validator is used to validate each item in the collection.
//
// Parameters:
// - itemValidator: The validator used to validate each item in the collection.
//
// Returns:
// - *Collection[T]: A pointer to the newly created Collection.
func NewCollection[T any](itemValidator contract.Validator[T]) *Collection[T] {
	return &Collection[T]{itemValidator: itemValidator}
}

// Validate checks if each item in the collection is valid using the item validator.
// It returns an error if any item is invalid. The errors are aggregated into an Errors structure.
//
// Parameters:
// - items: The slice of items to be validated.
//
// Returns:
// - error: An error if any item in the collection is invalid, or nil if all items are valid.
func (c *Collection[T]) Validate(items []T) error {
	errs := &errors.Errors{}
	for i, item := range items {
		if err := c.itemValidator.Validate(item); err != nil {
			field := fmt.Sprintf("[%d]", i)
			if nestedErrs, ok := err.(*errors.Errors); ok {
				errs.AddNested(field, nestedErrs)
			} else {
				errs.Add(field, err.Error())
			}
		}
	}
	if errs.IsEmpty() {
		return nil
	}
	return errs
}
