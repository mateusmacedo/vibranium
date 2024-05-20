package validation

import "github.com/mateusmacedo/vibranium/validation/contract"

type Composite[T any] struct {
    validators []WithContext[T]
}

type WithContext[T any] struct {
    Context   string
    Validator contract.Validator[T]
}

func NewComposite[T any]() *Composite[T] {
    return &Composite[T]{validators: []WithContext[T]{}}
}

func (cv *Composite[T]) Add(context string, validator contract.Validator[T]) {
    cv.validators = append(cv.validators, WithContext[T]{Context: context, Validator: validator})
}

func (cv *Composite[T]) Validate(value T) error {
    errors := &Errors{}
    for _, vc := range cv.validators {
        if err := vc.Validator.Validate(value); err != nil {
            errors.Add(vc.Context, err)
        }
    }
    if errors.IsEmpty() {
        return nil
    }
    return errors
}
