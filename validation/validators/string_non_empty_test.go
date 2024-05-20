package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringNonEmpty(t *testing.T) {
	tests := []struct {
		name          string
		value         string
		expectedError string
	}{
		{
			name:          "Non-empty string",
			value:         "hello",
			expectedError: "",
		},
		{
			name:          "Empty string",
			value:         "",
			expectedError: "value cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := &StringNonEmpty{}
			err := validator.Validate(tt.value)
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
		})
	}
}
