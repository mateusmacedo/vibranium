package errors

import "strings"

type Error struct {
    Context string
    Err     error
}

func (e *Error) Error() string {
    return e.Context + ": " + e.Err.Error()
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
