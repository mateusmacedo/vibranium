package validation

import (
	"github.com/mateusmacedo/vibranium/validation/contract"
	"github.com/mateusmacedo/vibranium/validation/errors"
)

type Composite[T any] struct {
    validators []WithField[T]
}

type WithField[T any] struct {
    Field     string
    Validator contract.Validator[T]
}

func NewComposite[T any]() *Composite[T] {
    return &Composite[T]{validators: []WithField[T]{}}
}

func (c *Composite[T]) AddValidator(field string, validator contract.Validator[T]) {
    c.validators = append(c.validators, WithField[T]{Field: field, Validator: validator})
}

func (c *Composite[T]) Validate(value T) error {
    errs := &errors.Errors{}
    for _, vc := range c.validators {
        if err := vc.Validator.Validate(value); err != nil {
            errs.Add(vc.Field, err)
        }
    }
    if errs.IsEmpty() {
        return nil
    }
    return errs
}
