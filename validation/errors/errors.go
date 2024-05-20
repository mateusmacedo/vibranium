package errors

import "strings"

// Error represents a validation error with a specific field and an error message.
type Error struct {
	Field string // Field is the name of the field that caused the validation error.
	Err   string // Err is the error message describing the validation failure.
}

// Errors aggregates multiple validation errors.
type Errors struct {
	List []Error // List contains all the validation errors.
}

// Error returns a string representation of all the validation errors.
// Each error is formatted as "Field: ErrorMessage" and errors are separated by newlines.
//
// Returns:
// - string: A string representation of all validation errors.
func (e *Errors) Error() string {
	var errorMessages []string
	for _, err := range e.List {
		errorMessages = append(errorMessages, err.Field+": "+err.Err)
	}
	return strings.Join(errorMessages, "\n")
}

// Add appends a new validation error to the list.
// It takes the field name and error message as arguments.
//
// Parameters:
// - field: The name of the field that caused the validation error.
// - message: The error message describing the validation failure.
func (e *Errors) Add(field, message string) {
	e.List = append(e.List, Error{Field: field, Err: message})
}

// AddNested appends nested validation errors to the list.
// It takes a field name and a pointer to another Errors instance.
// The field name is prefixed to each nested error's field name.
//
// Parameters:
// - field: The name of the field that caused the validation error.
// - nestedErrors: A pointer to another Errors instance containing nested validation errors.
func (e *Errors) AddNested(field string, nestedErrors *Errors) {
	for _, nestedErr := range nestedErrors.List {
		e.List = append(e.List, Error{Field: field + "." + nestedErr.Field, Err: nestedErr.Err})
	}
}

// IsEmpty checks if the list of validation errors is empty.
// It returns true if there are no errors, and false otherwise.
//
// Returns:
// - bool: True if there are no validation errors, false otherwise.
func (e *Errors) IsEmpty() bool {
	return len(e.List) == 0
}
