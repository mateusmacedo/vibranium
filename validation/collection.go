package validation

import (
	"fmt"

	"github.com/mateusmacedo/vibranium/validation/contract"
	"github.com/mateusmacedo/vibranium/validation/errors"
)

type Collection[T any] struct {
	itemValidator contract.Validator[T]
}

func NewCollection[T any](itemValidator contract.Validator[T]) *Collection[T] {
	return &Collection[T]{itemValidator: itemValidator}
}

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
