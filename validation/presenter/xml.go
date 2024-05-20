package presenter

import (
	"encoding/xml"

    "github.com/mateusmacedo/vibranium/validation/errors"
)

type XMLPresenter struct{}

func (p *XMLPresenter) Present(errors *errors.Errors) string {
    type XMLError struct {
        Field  string   `xml:"field"`
        Errors []string `xml:"error"`
    }

    errorMap := make(map[string][]string)
    for _, err := range errors.List {
        field := err.Field
        errorMap[field] = append(errorMap[field], err.Err.Error())
    }

    var result []XMLError
    for field, errs := range errorMap {
        result = append(result, XMLError{
            Field:  field,
            Errors: errs,
        })
    }

    xmlData, _ := xml.MarshalIndent(result, "", "  ")
    return string(xmlData)
}
