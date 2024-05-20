package validation

import "github.com/mateusmacedo/vibranium/validation/contract"

type Builder[T any] struct {
    composite *Composite[T]
}

func NewBuilder[T any]() *Builder[T] {
    return &Builder[T]{composite: NewComposite[T]()}
}

func (b *Builder[T]) Add(validator contract.Validator[T]) *Builder[T] {
    b.composite.Add(validator)
    return b
}

func (b *Builder[T]) Build() *Composite[T] {
    return b.composite
}
