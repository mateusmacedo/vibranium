package validation

import (
	"errors" // Using the standard errors package for error generation
	"testing"

	"github.com/stretchr/testify/assert"

	customErrors "github.com/mateusmacedo/vibranium/validation/errors"
	"github.com/mateusmacedo/vibranium/validation/mocks"
)

func TestBuilder(t *testing.T) {
	type User struct {
		Name  string
		Email string
	}

	tests := []struct {
		name           string
		initialBuilder *Builder[User]
		operations     func(*Builder[User])
		value          User
		expectedError  *customErrors.Errors
	}{
		{
			name:           "No validators",
			initialBuilder: NewBuilder[User](),
			operations:     func(b *Builder[User]) {},
			value:          User{Name: "John", Email: "john@example.com"},
			expectedError:  nil,
		},
		{
			name:           "Single validator passes",
			initialBuilder: NewBuilder[User](),
			operations: func(b *Builder[User]) {
				mockValidator := mocks.NewMockValidator[User](t)
				mockValidator.EXPECT().Validate(User{Name: "John", Email: "john@example.com"}).Return(nil)
				b.Add("Name", mockValidator)
			},
			value:         User{Name: "John", Email: "john@example.com"},
			expectedError: nil,
		},
		{
			name:           "Single validator fails",
			initialBuilder: NewBuilder[User](),
			operations: func(b *Builder[User]) {
				mockValidator := mocks.NewMockValidator[User](t)
				mockValidator.EXPECT().Validate(User{Name: "John", Email: "john@example.com"}).Return(errors.New("value cannot be empty"))
				b.Add("Name", mockValidator)
			},
			value:         User{Name: "John", Email: "john@example.com"},
			expectedError: &customErrors.Errors{List: []customErrors.Error{{Field: "Name", Err: "value cannot be empty"}}},
		},
		{
			name:           "Multiple validators",
			initialBuilder: NewBuilder[User](),
			operations: func(b *Builder[User]) {
				mockNameValidator := mocks.NewMockValidator[User](t)
				mockNameValidator.EXPECT().Validate(User{Name: "John", Email: "john@example.com"}).Return(nil)

				mockEmailValidator := mocks.NewMockValidator[User](t)
				mockEmailValidator.EXPECT().Validate(User{Name: "John", Email: "john@example.com"}).Return(errors.New("invalid email"))

				b.Add("Name", mockNameValidator)
				b.Add("Email", mockEmailValidator)
			},
			value:         User{Name: "John", Email: "john@example.com"},
			expectedError: &customErrors.Errors{List: []customErrors.Error{{Field: "Email", Err: "invalid email"}}},
		},
		{
			name:           "Nested validator errors",
			initialBuilder: NewBuilder[User](),
			operations: func(b *Builder[User]) {
				mockEmailValidator := mocks.NewMockValidator[User](t)
				nestedErrors := &customErrors.Errors{}
				nestedErrors.Add("Domain", "invalid domain")
				mockEmailValidator.EXPECT().Validate(User{Name: "John", Email: "john@example.com"}).Return(nestedErrors)
				b.Add("Email", mockEmailValidator)
			},
			value: User{Name: "John", Email: "john@example.com"},
			expectedError: &customErrors.Errors{List: []customErrors.Error{
				{Field: "Email.Domain", Err: "invalid domain"},
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.initialBuilder
			tt.operations(b)

			composite := b.Build()
			err := composite.Validate(tt.value)

			if tt.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tt.expectedError, err)
			}
		})
	}
}
