package presenter

import (
	"strings"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

// TextPresenter is responsible for presenting validation errors in plain text format.
type TextPresenter struct{}

// Present converts validation errors into a plain text formatted string.
// Each error is represented as "Field: ErrorMessage" and errors are separated by newlines.
//
// Parameters:
// - errors: A pointer to an Errors structure containing the validation errors.
//
// Returns:
// - string: A plain text formatted string representing the validation errors.
func (p *TextPresenter) Present(errors *errors.Errors) string {
	var result []string
	// Iterate over each error in the list and format it as "Field: ErrorMessage"
	for _, err := range errors.List {
		result = append(result, err.Field+": "+err.Err)
	}
	// Join all formatted errors into a single string separated by newlines
	return strings.Join(result, "\n")
}
