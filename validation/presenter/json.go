package presenter

import (
	"encoding/json"
	"strings"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

// JSONPresenter is responsible for presenting validation errors in JSON format.
type JSONPresenter struct{}

// Present converts validation errors into a JSON formatted string.
// It organizes the errors into a nested JSON structure based on the field names.
//
// Parameters:
// - errors: A pointer to an Errors structure containing the validation errors.
//
// Returns:
// - string: A JSON formatted string representing the validation errors.
func (p *JSONPresenter) Present(errors *errors.Errors) string {
	result := make(map[string]interface{})

	// Iterate over each error in the list
	for _, err := range errors.List {
		fields := strings.Split(err.Field, ".") // Split the field name by '.'
		current := result
		for i, field := range fields {
			if i == len(fields)-1 {
				// If it's the last field, append the error message to the list
				if current[field] == nil {
					current[field] = []string{}
				}
				current[field] = append(current[field].([]string), err.Err)
			} else {
				// If it's not the last field, traverse or create nested maps
				if current[field] == nil {
					current[field] = make(map[string]interface{})
				}
				current = current[field].(map[string]interface{})
			}
		}
	}

	// Marshal the result map into a pretty-printed JSON string
	jsonData, _ := json.MarshalIndent(result, "", "  ")
	return string(jsonData)
}
