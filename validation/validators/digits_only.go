package validators

import (
	"errors"
	"unicode"
)

type DigitsOnly struct{}

func (v DigitsOnly) Validate(value string) error {
    for _, char := range value {
        if !unicode.IsDigit(char) {
            return errors.New("value must contain only digits")
        }
    }
    return nil
}
