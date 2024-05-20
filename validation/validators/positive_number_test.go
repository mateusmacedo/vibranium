package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositiveNumber(t *testing.T) {
	tests := []struct {
		name          string
		value         int
		expectedError string
	}{
		{
			name:          "Positive number",
			value:         10,
			expectedError: "",
		},
		{
			name:          "Zero",
			value:         0,
			expectedError: "value must be positive",
		},
		{
			name:          "Negative number",
			value:         -5,
			expectedError: "value must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := &PositiveNumber{}
			err := validator.Validate(tt.value)
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
		})
	}
}
