package presenter

import (
	"strings"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

type TextPresenter struct{}

func (p *TextPresenter) Present(errors *errors.Errors) string {
	var result []string
	for _, err := range errors.List {
		result = append(result, err.Field+": "+err.Err)
	}
	return strings.Join(result, "\n")
}
