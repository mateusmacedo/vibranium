package presenter

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

func TestXMLPresenter(t *testing.T) {
	tests := []struct {
		name          string
		initialErrors *errors.Errors
		expectedOutput string
	}{
		{
			name: "No errors",
			initialErrors: &errors.Errors{},
			expectedOutput: "<XMLError></XMLError>",
		},
		{
			name: "Single error",
			initialErrors: &errors.Errors{
				List: []errors.Error{
					{Field: "Name", Err: "value cannot be empty"},
				},
			},
			expectedOutput: `<XMLError><Name><error>value cannot be empty</error></Name></XMLError>`,
		},
		{
			name: "Multiple errors",
			initialErrors: &errors.Errors{
				List: []errors.Error{
					{Field: "Name", Err: "value cannot be empty"},
					{Field: "Age", Err: "value must be positive"},
				},
			},
			expectedOutput: `<XMLError><Name><error>value cannot be empty</error></Name><Age><error>value must be positive</error></Age></XMLError>`,
		},
		{
			name: "Nested errors",
			initialErrors: &errors.Errors{
				List: []errors.Error{
					{Field: "Address.Street", Err: "value cannot be empty"},
					{Field: "Address.Zip", Err: "value must be positive"},
				},
			},
			expectedOutput: `<XMLError><Address><Street><error>value cannot be empty</error></Street><Zip><error>value must be positive</error></Zip></Address></XMLError>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			presenter := &XMLPresenter{}

			output := presenter.Present(tt.initialErrors)

			assertXMLEqual(t, tt.expectedOutput, output)
		})
	}
}

func assertXMLEqual(t *testing.T, expected, actual string) {
	var expectedXML, actualXML interface{}
	xml.Unmarshal([]byte(expected), &expectedXML)
	xml.Unmarshal([]byte(actual), &actualXML)

	assert.Equal(t, expectedXML, actualXML)
}
