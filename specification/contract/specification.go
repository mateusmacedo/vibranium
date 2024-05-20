package contract

// Candidate represents a generic type that can be used as a candidate for validation.
type Candidate any

// Specification is an interface that defines a contract for a specification.
// A specification is used to determine if a candidate object satisfies certain criteria.
type Specification[T Candidate] interface {
	// IsSatisfiedBy checks if the candidate object satisfies the specification.
	//
	// Parameters:
	// - candidate: The candidate object to be evaluated.
	//
	// Returns:
	// - bool: True if the candidate satisfies the specification, false otherwise.
	IsSatisfiedBy(candidate T) bool
}
