package validation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	customErrors "github.com/mateusmacedo/vibranium/validation/errors"
	"github.com/mateusmacedo/vibranium/validation/mocks"
)

func TestComposite(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name          string
		initialErrors *Composite[User]
		operations    func(*Composite[User])
		value         User
		expectedError *customErrors.Errors
	}{
		{
			name:          "No validators",
			initialErrors: NewComposite[User](),
			operations:    func(c *Composite[User]) {},
			value:         User{Name: "John", Age: 30},
			expectedError: nil,
		},
		{
			name:          "Single validator passes",
			initialErrors: NewComposite[User](),
			operations: func(c *Composite[User]) {
				mockValidator := mocks.NewMockValidator[User](t)
				mockValidator.EXPECT().Validate(User{Name: "John", Age: 30}).Return(nil)
				c.AddValidator("Name", mockValidator)
			},
			value:         User{Name: "John", Age: 30},
			expectedError: nil,
		},
		{
			name:          "Single validator fails",
			initialErrors: NewComposite[User](),
			operations: func(c *Composite[User]) {
				mockValidator := mocks.NewMockValidator[User](t)
				mockValidator.EXPECT().Validate(User{Name: "John", Age: 30}).Return(errors.New("value cannot be empty"))
				c.AddValidator("Name", mockValidator)
			},
			value:         User{Name: "John", Age: 30},
			expectedError: &customErrors.Errors{List: []customErrors.Error{{Field: "Name", Err: "value cannot be empty"}}},
		},
		{
			name:          "Multiple validators",
			initialErrors: NewComposite[User](),
			operations: func(c *Composite[User]) {
				mockNameValidator := mocks.NewMockValidator[User](t)
				mockNameValidator.EXPECT().Validate(User{Name: "John", Age: 30}).Return(nil)

				mockAgeValidator := mocks.NewMockValidator[User](t)
				mockAgeValidator.EXPECT().Validate(User{Name: "John", Age: 30}).Return(errors.New("value must be positive"))

				c.AddValidator("Name", mockNameValidator)
				c.AddValidator("Age", mockAgeValidator)
			},
			value:         User{Name: "John", Age: 30},
			expectedError: &customErrors.Errors{List: []customErrors.Error{{Field: "Age", Err: "value must be positive"}}},
		},
		{
			name:          "Nested validator errors",
			initialErrors: NewComposite[User](),
			operations: func(c *Composite[User]) {
				mockAddressValidator := mocks.NewMockValidator[User](t)
				nestedErrors := &customErrors.Errors{}
				nestedErrors.Add("Street", "value cannot be empty")
				nestedErrors.Add("Zip", "value must be positive")
				mockAddressValidator.EXPECT().Validate(User{Name: "John", Age: 30}).Return(nestedErrors)
				c.AddValidator("Address", mockAddressValidator)
			},
			value: User{Name: "John", Age: 30},
			expectedError: &customErrors.Errors{List: []customErrors.Error{
				{Field: "Address.Street", Err: "value cannot be empty"},
				{Field: "Address.Zip", Err: "value must be positive"},
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
