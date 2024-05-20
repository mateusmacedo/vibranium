// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	contract "github.com/mateusmacedo/vibranium/validation/contract"
	mock "github.com/stretchr/testify/mock"
)

// MockValidator is an autogenerated mock type for the Validator type
type MockValidator[T contract.Value] struct {
	mock.Mock
}

type MockValidator_Expecter[T contract.Value] struct {
	mock *mock.Mock
}

func (_m *MockValidator[T]) EXPECT() *MockValidator_Expecter[T] {
	return &MockValidator_Expecter[T]{mock: &_m.Mock}
}

// Validate provides a mock function with given fields: value
func (_m *MockValidator[T]) Validate(value T) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(T) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockValidator_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type MockValidator_Validate_Call[T contract.Value] struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - value T
func (_e *MockValidator_Expecter[T]) Validate(value interface{}) *MockValidator_Validate_Call[T] {
	return &MockValidator_Validate_Call[T]{Call: _e.mock.On("Validate", value)}
}

func (_c *MockValidator_Validate_Call[T]) Run(run func(value T)) *MockValidator_Validate_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(T))
	})
	return _c
}

func (_c *MockValidator_Validate_Call[T]) Return(_a0 error) *MockValidator_Validate_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockValidator_Validate_Call[T]) RunAndReturn(run func(T) error) *MockValidator_Validate_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewMockValidator creates a new instance of MockValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockValidator[T contract.Value](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockValidator[T] {
	mock := &MockValidator[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}