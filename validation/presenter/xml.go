package presenter

import (
	"encoding/xml"
	"strings"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

// XMLError represents a structured format for validation errors in XML.
type XMLError struct {
	XMLName xml.Name   // XMLName is the name of the XML element.
	Errors  []string   `xml:"error,omitempty"` // Errors contains the list of error messages.
	Nested  []XMLError `xml:",omitempty"`      // Nested contains nested XMLError elements for hierarchical errors.
}

// XMLPresenter is responsible for presenting validation errors in XML format.
type XMLPresenter struct{}

// Present converts validation errors into an XML formatted string.
// It organizes the errors into a nested XML structure based on the field names.
//
// Parameters:
// - errs: A pointer to an Errors structure containing the validation errors.
//
// Returns:
// - string: An XML formatted string representing the validation errors.
func (p *XMLPresenter) Present(errs *errors.Errors) string {
	result := make(map[string]interface{})

	// Iterate over each error in the list
	for _, err := range errs.List {
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

	// Convert the nested map structure to XMLError
	xmlErrors := convertToXMLError(result)
	// Marshal the XMLError structure into a pretty-printed XML string
	xmlData, _ := xml.MarshalIndent(xmlErrors, "", "  ")
	return string(xmlData)
}

// convertToXMLError recursively converts a nested map structure to an XMLError structure.
//
// Parameters:
// - data: A nested map structure representing the errors.
//
// Returns:
// - XMLError: An XMLError structure representing the nested errors.
func convertToXMLError(data map[string]interface{}) XMLError {
	var result XMLError
	for key, value := range data {
		if nested, ok := value.(map[string]interface{}); ok {
			// If the value is a nested map, recursively convert it to XMLError
			nestedErrors := convertToXMLError(nested)
			nestedErrors.XMLName = xml.Name{Local: key}
			result.Nested = append(result.Nested, nestedErrors)
		} else if errors, ok := value.([]string); ok {
			// If the value is a slice of error strings, set them in the Errors field
			result.XMLName = xml.Name{Local: key}
			result.Errors = errors
		}
	}
	return result
}
