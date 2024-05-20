package errors

import "strings"

type Error struct {
    Field string
    Err     error
}

func (e *Error) Error() string {
    return e.Field + ": " + e.Err.Error()
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

func (e *Errors) Add(field string, err error) {
    if nestedErrors, ok := err.(*Errors); ok {
        for _, nestedErr := range nestedErrors.List {
            e.List = append(e.List, Error{Field: field + "." + nestedErr.Field, Err: nestedErr.Err})
        }
    } else {
        e.List = append(e.List, Error{Field: field, Err: err})
    }
}

func (e *Errors) IsEmpty() bool {
    return len(e.List) == 0
}
