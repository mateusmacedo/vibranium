package errors

import (
	"testing"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		name          string
		initialErrors Errors
		operations    func(*Errors)
		expectedList  []Error
		expectedError string
		expectedEmpty bool
	}{
		{
			name:          "Empty errors",
			initialErrors: Errors{},
			operations:    func(e *Errors) {},
			expectedList:  []Error{},
			expectedError: "",
			expectedEmpty: true,
		},
		{
			name:          "Add single error",
			initialErrors: Errors{},
			operations: func(e *Errors) {
				e.Add("Name", "value cannot be empty")
			},
			expectedList:  []Error{{Field: "Name", Err: "value cannot be empty"}},
			expectedError: "Name: value cannot be empty",
			expectedEmpty: false,
		},
		{
			name:          "Add multiple errors",
			initialErrors: Errors{},
			operations: func(e *Errors) {
				e.Add("Name", "value cannot be empty")
				e.Add("Age", "value must be positive")
			},
			expectedList: []Error{
				{Field: "Name", Err: "value cannot be empty"},
				{Field: "Age", Err: "value must be positive"},
			},
			expectedError: "Name: value cannot be empty\nAge: value must be positive",
			expectedEmpty: false,
		},
		{
			name:          "Add nested errors",
			initialErrors: Errors{},
			operations: func(e *Errors) {
				nestedErrors := &Errors{}
				nestedErrors.Add("Street", "value cannot be empty")
				nestedErrors.Add("Zip", "value must be positive")
				e.AddNested("Address", nestedErrors)
			},
			expectedList: []Error{
				{Field: "Address.Street", Err: "value cannot be empty"},
				{Field: "Address.Zip", Err: "value must be positive"},
			},
			expectedError: "Address.Street: value cannot be empty\nAddress.Zip: value must be positive",
			expectedEmpty: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &tt.initialErrors
			tt.operations(e)

			if !equalErrors(e.List, tt.expectedList) {
				t.Errorf("expected errors list %v, got %v", tt.expectedList, e.List)
			}
			if gotError := e.Error(); gotError != tt.expectedError {
				t.Errorf("expected error string %q, got %q", tt.expectedError, gotError)
			}
			if gotEmpty := e.IsEmpty(); gotEmpty != tt.expectedEmpty {
				t.Errorf("expected IsEmpty %v, got %v", tt.expectedEmpty, gotEmpty)
			}
		})
	}
}

func equalErrors(a, b []Error) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
