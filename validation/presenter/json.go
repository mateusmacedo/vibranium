package presenter

import (
	"encoding/json"
	"strings"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

type JSONPresenter struct{}

func (p *JSONPresenter) Present(errors *errors.Errors) string {
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

	jsonData, _ := json.MarshalIndent(result, "", "  ")
	return string(jsonData)
}
