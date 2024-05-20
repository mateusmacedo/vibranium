package errors

import "strings"

type Error struct {
	Field string
	Err   string
}

type Errors struct {
	List []Error
}

func (e *Errors) Error() string {
	var errorMessages []string
	for _, err := range e.List {
		errorMessages = append(errorMessages, err.Field+": "+err.Err)
	}
	return strings.Join(errorMessages, "\n")
}

func (e *Errors) Add(field, message string) {
	e.List = append(e.List, Error{Field: field, Err: message})
}

func (e *Errors) AddNested(field string, nestedErrors *Errors) {
	for _, nestedErr := range nestedErrors.List {
		e.List = append(e.List, Error{Field: field + "." + nestedErr.Field, Err: nestedErr.Err})
	}
}

func (e *Errors) IsEmpty() bool {
	return len(e.List) == 0
}
