package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitsOnly(t *testing.T) {
	tests := []struct {
		name          string
		value         string
		expectedError string
	}{
		{
			name:          "Valid digits",
			value:         "123456",
			expectedError: "",
		},
		{
			name:          "Contains letters",
			value:         "123a56",
			expectedError: "value must contain only digits",
		},
		{
			name:          "Contains special characters",
			value:         "123@56",
			expectedError: "value must contain only digits",
		},
		{
			name:          "Empty string",
			value:         "",
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := &DigitsOnly{}
			err := validator.Validate(tt.value)
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
		})
	}
}
