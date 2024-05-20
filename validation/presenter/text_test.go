package presenter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/vibranium/validation/errors"
)

func TestTextPresenter(t *testing.T) {
	tests := []struct {
		name          string
		initialErrors *errors.Errors
		expectedOutput string
	}{
		{
			name: "No errors",
			initialErrors: &errors.Errors{},
			expectedOutput: "",
		},
		{
			name: "Single error",
			initialErrors: &errors.Errors{
				List: []errors.Error{
					{Field: "Name", Err: "value cannot be empty"},
				},
			},
			expectedOutput: "Name: value cannot be empty",
		},
		{
			name: "Multiple errors",
			initialErrors: &errors.Errors{
				List: []errors.Error{
					{Field: "Name", Err: "value cannot be empty"},
					{Field: "Age", Err: "value must be positive"},
				},
			},
			expectedOutput: "Name: value cannot be empty\nAge: value must be positive",
		},
		{
			name: "Nested errors",
			initialErrors: &errors.Errors{
				List: []errors.Error{
					{Field: "Address.Street", Err: "value cannot be empty"},
					{Field: "Address.Zip", Err: "value must be positive"},
				},
			},
			expectedOutput: "Address.Street: value cannot be empty\nAddress.Zip: value must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			presenter := &TextPresenter{}

			output := presenter.Present(tt.initialErrors)

			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
