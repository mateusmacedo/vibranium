package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExactLength(t *testing.T) {
	tests := []struct {
		name          string
		value         string
		length        int
		expectedError string
	}{
		{
			name:          "Valid length",
			value:         "12345",
			length:        5,
			expectedError: "",
		},
		{
			name:          "Too short",
			value:         "123",
			length:        5,
			expectedError: "value must be exactly 5 characters long",
		},
		{
			name:          "Too long",
			value:         "1234567",
			length:        5,
			expectedError: "value must be exactly 5 characters long",
		},
		{
			name:          "Empty string",
			value:         "",
			length:        0,
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := &ExactLength{Length: tt.length}
			err := validator.Validate(tt.value)
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
		})
	}
}
