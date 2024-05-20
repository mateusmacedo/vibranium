package validation

import "github.com/mateusmacedo/vibranium/validation/contract"

type Composite[T any] struct {
    validators []contract.Validator[T]
}

func NewComposite[T any]() *Composite[T] {
    return &Composite[T]{validators: []contract.Validator[T]{}}
}

func (c *Composite[T]) Add(validator contract.Validator[T]) *Composite[T] {
    c.validators = append(c.validators, validator)
    return c
}

func (c *Composite[T]) Validate(value T) error {
    errors := &Errors{}
    for _, validator := range c.validators {
        if err := validator.Validate(value); err != nil {
            errors.Add(err)
        }
    }
    if errors.IsEmpty() {
        return nil
    }
    return errors
}
