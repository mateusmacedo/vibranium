package specification

import "github.com/mateusmacedo/vibranium/specification/contract"

// OrSpecification represents a logical OR operation between multiple specifications.
type OrSpecification[T contract.Candidate] struct {
	specs []contract.Specification[T] // specs is a slice of specifications to be combined using OR.
}

// NewOrSpecification creates a new OrSpecification that represents a logical OR operation between multiple specifications.
// It takes a variadic parameter `specs` of type `contract.Specification[T]` which represents the specifications to be combined.
// It returns a pointer to the created `OrSpecification[T]`.
//
// Parameters:
// - specs: A variadic parameter of specifications to be combined using OR.
//
// Returns:
// - *OrSpecification[T]: A pointer to the newly created OrSpecification.
func NewOrSpecification[T contract.Candidate](specs ...contract.Specification[T]) *OrSpecification[T] {
	return &OrSpecification[T]{specs: specs}
}

// IsSatisfiedBy checks if the given candidate satisfies any of the specifications in the OrSpecification.
// It iterates over each specification and returns true if any of them is satisfied by the candidate.
// If none of the specifications are satisfied, it returns false.
//
// Parameters:
// - candidate: The candidate object to be evaluated.
//
// Returns:
// - bool: True if the candidate satisfies any of the specifications, false otherwise.
func (s *OrSpecification[T]) IsSatisfiedBy(candidate T) bool {
	for _, spec := range s.specs {
		if spec.IsSatisfiedBy(candidate) {
			return true
		}
	}
	return false
}
