package validators

import "errors"

type StringNonEmpty struct{}

func (v StringNonEmpty) Validate(value string) error {
    if value == "" {
        return errors.New("value cannot be empty")
    }
    return nil
}
