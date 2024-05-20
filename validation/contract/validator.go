package contract

type Value any

type Validator[T Value] interface {
    Validate(value T) error
}

type ValidatorFunc[T Value] func(value T) error

func (f ValidatorFunc[T]) Validate(value T) error {
    return f(value)
}
