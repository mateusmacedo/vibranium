// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	contract "github.com/mateusmacedo/vibranium/specification/contract"
	mock "github.com/stretchr/testify/mock"
)

// MockSpecification is an autogenerated mock type for the Specification type
type MockSpecification[T contract.Candidate] struct {
	mock.Mock
}

type MockSpecification_Expecter[T contract.Candidate] struct {
	mock *mock.Mock
}

func (_m *MockSpecification[T]) EXPECT() *MockSpecification_Expecter[T] {
	return &MockSpecification_Expecter[T]{mock: &_m.Mock}
}

// IsSatisfiedBy provides a mock function with given fields: candidate
func (_m *MockSpecification[T]) IsSatisfiedBy(candidate T) bool {
	ret := _m.Called(candidate)

	if len(ret) == 0 {
		panic("no return value specified for IsSatisfiedBy")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(T) bool); ok {
		r0 = rf(candidate)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockSpecification_IsSatisfiedBy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSatisfiedBy'
type MockSpecification_IsSatisfiedBy_Call[T contract.Candidate] struct {
	*mock.Call
}

// IsSatisfiedBy is a helper method to define mock.On call
//   - candidate T
func (_e *MockSpecification_Expecter[T]) IsSatisfiedBy(candidate interface{}) *MockSpecification_IsSatisfiedBy_Call[T] {
	return &MockSpecification_IsSatisfiedBy_Call[T]{Call: _e.mock.On("IsSatisfiedBy", candidate)}
}

func (_c *MockSpecification_IsSatisfiedBy_Call[T]) Run(run func(candidate T)) *MockSpecification_IsSatisfiedBy_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(T))
	})
	return _c
}

func (_c *MockSpecification_IsSatisfiedBy_Call[T]) Return(_a0 bool) *MockSpecification_IsSatisfiedBy_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSpecification_IsSatisfiedBy_Call[T]) RunAndReturn(run func(T) bool) *MockSpecification_IsSatisfiedBy_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewMockSpecification creates a new instance of MockSpecification. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSpecification[T contract.Candidate](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSpecification[T] {
	mock := &MockSpecification[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
