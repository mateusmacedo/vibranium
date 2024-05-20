package presenter

import (
	"encoding/json"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

type JSONPresenter struct{}

func (p *JSONPresenter) Present(errors *errors.Errors) string {
	type ErrorMessage struct {
		Field string   `json:"field"`
		Errors  []string `json:"errors"`
	}

	errorMap := make(map[string][]string)
	for _, err := range errors.List {
		field := err.Field
		errorMap[field] = append(errorMap[field], err.Err.Error())
	}

	var result []ErrorMessage
	for field, errs := range errorMap {
		result = append(result, ErrorMessage{
			Field: field,
			Errors:  errs,
		})
	}

	jsonData, _ := json.Marshal(result)
	return string(jsonData)
}
