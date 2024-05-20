package presenter

import (
	"encoding/json"
	"strings"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

type JSONPresenter struct{}

func (p *JSONPresenter) Present(errors *errors.Errors) string {
    type ErrorMessage struct {
        Context string `json:"context"`
        Error   string `json:"error"`
    }

    errorMap := make(map[string][]string)
    for _, err := range errors.List {
        context := err.Context
        errorMap[context] = append(errorMap[context], err.Err.Error())
    }

    var result []ErrorMessage
    for context, errs := range errorMap {
        result = append(result, ErrorMessage{
            Context: context,
            Error:   strings.Join(errs, ";\n"),
        })
    }

    jsonData, _ := json.Marshal(result)
    return string(jsonData)
}
