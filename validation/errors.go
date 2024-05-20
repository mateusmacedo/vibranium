package validation

import (
	"strings"
)

type Errors struct {
    errors []error
}

func (e *Errors) Error() string {
    var errorMessages []string
    for _, err := range e.errors {
        errorMessages = append(errorMessages, err.Error())
    }
    return strings.Join(errorMessages, ";\n")
}

func (e *Errors) Add(err error) {
    e.errors = append(e.errors, err)
}

func (e *Errors) IsEmpty() bool {
    return len(e.errors) == 0
}
