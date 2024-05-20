package validation

import "github.com/mateusmacedo/vibranium/validation/contract"

type Builder[T any] struct {
    compositeValidator *Composite[T]
}

func NewBuilder[T any]() *Builder[T] {
    return &Builder[T]{compositeValidator: NewComposite[T]()}
}

func (vb *Builder[T]) Add(context string, validator contract.Validator[T]) *Builder[T] {
    vb.compositeValidator.Add(context, validator)
    return vb
}

func (vb *Builder[T]) Build() *Composite[T] {
    return vb.compositeValidator
}
