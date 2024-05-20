package validators

import (
	"errors"
	"fmt"
)

type ExactLength struct {
    Length int
}

func (v ExactLength) Validate(value string) error {
    if len(value) != v.Length {
        return errors.New(fmt.Sprintf("value must be exactly %d characters long", v.Length))
    }
    return nil
}
