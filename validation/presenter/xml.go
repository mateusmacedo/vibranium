package presenter

import (
	"encoding/xml"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

type XMLPresenter struct{}

func (p *XMLPresenter) Present(errors *errors.Errors) string {
    type XMLError struct {
        Context string `xml:"context"`
        Error   string `xml:"error"`
    }
    var result []XMLError
    for _, err := range errors.List {
        result = append(result, XMLError{
            Context: err.Context,
            Error:   err.Err.Error(),
        })
    }
    xmlData, _ := xml.MarshalIndent(result, "", "  ")
    return string(xmlData)
}
