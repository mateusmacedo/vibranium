package specification

import "github.com/mateusmacedo/vibranium/specification/contract"

// AndSpecification represents a composite specification that combines multiple specifications using a logical AND operation.
type AndSpecification[T contract.Candidate] struct {
    specs []contract.Specification[T]
}

// NewAndSpecification creates a new AndSpecification that combines multiple specifications into one.
// It takes a variadic parameter `specs` of type `contract.Specification[T]` and returns a pointer to the created `AndSpecification[T]`.
func NewAndSpecification[T contract.Candidate](specs ...contract.Specification[T]) *AndSpecification[T] {
    return &AndSpecification[T]{specs: specs}
}

// IsSatisfiedBy checks if the given candidate satisfies all the specifications in the AndSpecification.
// It iterates over each specification and returns false if any of them is not satisfied.
// If all specifications are satisfied, it returns true.
func (s *AndSpecification[T]) IsSatisfiedBy(candidate T) bool {
    for _, spec := range s.specs {
        if !spec.IsSatisfiedBy(candidate) {
            return false
        }
    }
    return true
}