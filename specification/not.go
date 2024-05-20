package specification

import "github.com/mateusmacedo/vibranium/specification/contract"

// NotSpecification represents a specification that negates the result of another specification.
type NotSpecification[T contract.Candidate] struct {
	spec contract.Specification[T] // spec is the underlying specification to be negated.
}

// NewNotSpecification creates a new instance of the NotSpecification.
// It takes a specification as input and returns a pointer to the NotSpecification.
//
// Parameters:
// - spec: The specification to be negated.
//
// Returns:
// - *NotSpecification[T]: A pointer to the newly created NotSpecification.
func NewNotSpecification[T contract.Candidate](spec contract.Specification[T]) *NotSpecification[T] {
	return &NotSpecification[T]{spec: spec}
}

// IsSatisfiedBy checks if the given candidate satisfies the specification.
// It negates the result of the underlying specification's IsSatisfiedBy method.
//
// Parameters:
// - candidate: The candidate object to be evaluated.
//
// Returns:
// - bool: True if the candidate does not satisfy the underlying specification, false otherwise.
func (s *NotSpecification[T]) IsSatisfiedBy(candidate T) bool {
	return !s.spec.IsSatisfiedBy(candidate)
}
