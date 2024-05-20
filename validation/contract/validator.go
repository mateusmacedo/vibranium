package contract

type Value any

type Validator[T Value] interface {
    Validate(value T) error
}

type ValidationFunc[T Value] func(value T) error

func (f ValidationFunc[T]) Validate(value T) error {
    return f(value)
}
