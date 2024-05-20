package validation

import (
	"fmt"
	"strings"
)

type Error struct {
    Context string
    Err     error
}

func (ve *Error) Error() string {
    return fmt.Sprintf("%s: %s", ve.Context, ve.Err.Error())
}

type Errors struct {
    List []Error
}

func (e *Errors) Error() string {
    var errorMessages []string
    for _, err := range e.List {
        errorMessages = append(errorMessages, err.Error())
    }
    return strings.Join(errorMessages, ";\n")
}

func (e *Errors) Add(context string, err error) {
    e.List = append(e.List, Error{Context: context, Err: err})
}

func (e *Errors) IsEmpty() bool {
    return len(e.List) == 0
}

func IndentErrorMessages(errorMessages string, indent string) string {
    lines := strings.Split(errorMessages, ";\n")
    for i, line := range lines {
        lines[i] = indent + line
    }
    return strings.Join(lines, ";\n")
}
