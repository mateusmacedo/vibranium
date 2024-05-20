package validation

import (
	"fmt"

	"github.com/mateusmacedo/vibranium/validation/contract"
)

type Collection[T any] struct {
    itemValidator contract.Validator[T]
}

func NewCollection[T any](itemValidator contract.Validator[T]) *Collection[T] {
    return &Collection[T]{itemValidator: itemValidator}
}

func (cv *Collection[T]) Validate(items []T) error {
    errors := &Errors{}
    for i, item := range items {
        if err := cv.itemValidator.Validate(item); err != nil {
            context := fmt.Sprintf("Item %d", i)
            indentedError := IndentErrorMessages(err.Error(), "  ")
            errors.Add(context, fmt.Errorf("%s", indentedError))
        }
    }
    if errors.IsEmpty() {
        return nil
    }
    return errors
}
