package specification

import "github.com/mateusmacedo/vibranium/specification/contract"

// OrSpecification represents a logical OR operation between multiple specifications.
type OrSpecification[T contract.Candidate] struct {
    specs []contract.Specification[T]
}

// NewOrSpecification creates a new OrSpecification that represents a logical OR operation between multiple specifications.
// It takes a variadic parameter `specs` of type `contract.Specification[T]` which represents the specifications to be combined.
// It returns a pointer to the created `OrSpecification[T]`.
func NewOrSpecification[T contract.Candidate](specs ...contract.Specification[T]) *OrSpecification[T] {
    return &OrSpecification[T]{specs: specs}
}

// IsSatisfiedBy checks if the given candidate satisfies any of the specifications in the OrSpecification.
// It iterates over each specification and returns true if any of them is satisfied by the candidate.
// If none of the specifications are satisfied, it returns false.
func (s *OrSpecification[T]) IsSatisfiedBy(candidate T) bool {
    for _, spec := range s.specs {
        if spec.IsSatisfiedBy(candidate) {
            return true
        }
    }
    return false
}