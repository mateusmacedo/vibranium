package specification

import "github.com/mateusmacedo/vibranium/specification/contract"

// NotSpecification represents a specification that negates the result of another specification.
type NotSpecification[T contract.Candidate] struct {
    spec contract.Specification[T]
}

// NewNotSpecification creates a new instance of the NotSpecification.
// It takes a specification as input and returns a pointer to the NotSpecification.
func NewNotSpecification[T contract.Candidate](spec contract.Specification[T]) *NotSpecification[T] {
    return &NotSpecification[T]{spec: spec}
}

// IsSatisfiedBy checks if the given candidate satisfies the specification.
// It negates the result of the underlying specification's IsSatisfiedBy method.
func (s *NotSpecification[T]) IsSatisfiedBy(candidate T) bool {
    return !s.spec.IsSatisfiedBy(candidate)
}