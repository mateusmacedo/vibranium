package validators

import "errors"

type PositiveNumber struct{}

func (v PositiveNumber) Validate(value int) error {
    if value <= 0 {
        return errors.New("value must be positive")
    }
    return nil
}
