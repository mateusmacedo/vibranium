package validation

import "github.com/mateusmacedo/vibranium/validation/contract"

type Collection[T any] struct {
    itemValidator contract.Validator[T]
}

func NewCollection[T any](itemValidator contract.Validator[T]) *Collection[T] {
    return &Collection[T]{itemValidator: itemValidator}
}

func (c *Collection[T]) Validate(items []T) error {
    errors := &Errors{}
    for _, item := range items {
        if err := c.itemValidator.Validate(item); err != nil {
            errors.Add(err)
        }
    }
    if errors.IsEmpty() {
        return nil
    }
    return errors
}
