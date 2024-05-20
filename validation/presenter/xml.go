package presenter

import (
	"encoding/xml"
	"strings"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

type XMLPresenter struct{}

func (p *XMLPresenter) Present(errors *errors.Errors) string {
    type XMLError struct {
        Context string `xml:"context"`
        Error   string `xml:"error"`
    }

    errorMap := make(map[string][]string)
    for _, err := range errors.List {
        context := err.Context
        errorMap[context] = append(errorMap[context], err.Err.Error())
    }

    var result []XMLError
    for context, errs := range errorMap {
        result = append(result, XMLError{
            Context: context,
            Error:   strings.Join(errs, ";\n"),
        })
    }

    xmlData, _ := xml.MarshalIndent(result, "", "  ")
    return string(xmlData)
}
