package validation

import (
	"errors" // Using the standard errors package for error generation
	"testing"

	"github.com/stretchr/testify/assert"

	customErrors "github.com/mateusmacedo/vibranium/validation/errors"
	"github.com/mateusmacedo/vibranium/validation/mocks"
)

func TestCollection(t *testing.T) {
	type Item struct {
		Value string
	}

	tests := []struct {
		name          string
		initialErrors *Collection[Item]
		operations    func(*Collection[Item])
		value         []Item
		expectedError *customErrors.Errors
	}{
		{
			name:          "No items",
			initialErrors: NewCollection[Item](nil),
			operations:    func(c *Collection[Item]) {},
			value:         []Item{},
			expectedError: nil,
		},
		{
			name:          "Single item passes",
			initialErrors: NewCollection[Item](nil),
			operations: func(c *Collection[Item]) {
				mockValidator := mocks.NewMockValidator[Item](t)
				mockValidator.EXPECT().Validate(Item{Value: "valid"}).Return(nil)
				c.itemValidator = mockValidator
			},
			value:         []Item{{Value: "valid"}},
			expectedError: nil,
		},
		{
			name:          "Single item fails",
			initialErrors: NewCollection[Item](nil),
			operations: func(c *Collection[Item]) {
				mockValidator := mocks.NewMockValidator[Item](t)
				mockValidator.EXPECT().Validate(Item{Value: "invalid"}).Return(errors.New("value is invalid"))
				c.itemValidator = mockValidator
			},
			value:         []Item{{Value: "invalid"}},
			expectedError: &customErrors.Errors{List: []customErrors.Error{{Field: "[0]", Err: "value is invalid"}}},
		},
		{
			name:          "Multiple items",
			initialErrors: NewCollection[Item](nil),
			operations: func(c *Collection[Item]) {
				mockValidator := mocks.NewMockValidator[Item](t)
				mockValidator.EXPECT().Validate(Item{Value: "valid"}).Return(nil)
				mockValidator.EXPECT().Validate(Item{Value: "invalid"}).Return(errors.New("value is invalid"))
				c.itemValidator = mockValidator
			},
			value:         []Item{{Value: "valid"}, {Value: "invalid"}},
			expectedError: &customErrors.Errors{List: []customErrors.Error{{Field: "[1]", Err: "value is invalid"}}},
		},
		{
			name:          "Nested item errors",
			initialErrors: NewCollection[Item](nil),
			operations: func(c *Collection[Item]) {
				mockValidator := mocks.NewMockValidator[Item](t)
				nestedErrors := &customErrors.Errors{}
				nestedErrors.Add("SubField", "value cannot be empty")
				mockValidator.EXPECT().Validate(Item{Value: "nested"}).Return(nestedErrors)
				c.itemValidator = mockValidator
			},
			value: []Item{{Value: "nested"}},
			expectedError: &customErrors.Errors{List: []customErrors.Error{
				{Field: "[0].SubField", Err: "value cannot be empty"},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initialErrors
			tt.operations(c)

			err := c.Validate(tt.value)

			if tt.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tt.expectedError, err)
			}
		})
	}
}
