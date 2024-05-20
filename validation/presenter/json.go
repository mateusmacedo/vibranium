package presenter

import (
	"encoding/json"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

type JSONPresenter struct{}

func (p *JSONPresenter) Present(errors *errors.Errors) string {
    var result []map[string]string
    for _, err := range errors.List {
        result = append(result, map[string]string{
            "context": err.Context,
            "error":   err.Err.Error(),
        })
    }
    jsonData, _ := json.Marshal(result)
    return string(jsonData)
}
