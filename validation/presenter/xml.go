package presenter

import (
	"encoding/xml"
	"strings"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

type XMLError struct {
	XMLName xml.Name
	Errors  []string   `xml:"error,omitempty"`
	Nested  []XMLError `xml:",omitempty"`
}

type XMLPresenter struct{}

func (p *XMLPresenter) Present(errors *errors.Errors) string {
	result := make(map[string]interface{})

	for _, err := range errors.List {
		fields := strings.Split(err.Field, ".")
		current := result
		for i, field := range fields {
			if i == len(fields)-1 {
				if current[field] == nil {
					current[field] = []string{}
				}
				current[field] = append(current[field].([]string), err.Err)
			} else {
				if current[field] == nil {
					current[field] = make(map[string]interface{})
				}
				current = current[field].(map[string]interface{})
			}
		}
	}

	xmlErrors := convertToXMLError(result)
	xmlData, _ := xml.MarshalIndent(xmlErrors, "", "  ")
	return string(xmlData)
}

func convertToXMLError(data map[string]interface{}) XMLError {
	var result XMLError
	for key, value := range data {
		if nested, ok := value.(map[string]interface{}); ok {
			nestedErrors := convertToXMLError(nested)
			nestedErrors.XMLName = xml.Name{Local: key}
			result.Nested = append(result.Nested, nestedErrors)
		} else if errors, ok := value.([]string); ok {
			result.XMLName = xml.Name{Local: key}
			result.Errors = errors
		}
	}
	return result
}

func escapeXMLName(name string) string {
	return strings.ReplaceAll(strings.ReplaceAll(name, "[", ""), "]", "")
}
