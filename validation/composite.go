package validation

import (
	"github.com/mateusmacedo/vibranium/validation/contract"
	"github.com/mateusmacedo/vibranium/validation/errors"
)

type Composite[T any] struct {
    validators []ValidatorWithContext[T]
}

type ValidatorWithContext[T any] struct {
    Context   string
    Validator contract.Validator[T]
}

func NewComposite[T any]() *Composite[T] {
    return &Composite[T]{validators: []ValidatorWithContext[T]{}}
}

func (c *Composite[T]) AddValidator(context string, validator contract.Validator[T]) {
    c.validators = append(c.validators, ValidatorWithContext[T]{Context: context, Validator: validator})
}

func (c *Composite[T]) Validate(value T) error {
    errs := &errors.Errors{}
    for _, vc := range c.validators {
        if err := vc.Validator.Validate(value); err != nil {
            errs.Add(vc.Context, err)
        }
    }
    if errs.IsEmpty() {
        return nil
    }
    return errs
}
