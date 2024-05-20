package contract

// SpecificationBuilder is an interface for building specifications.
type SpecificationBuilder[T Candidate] interface {
	// WithSpecification adds a specification to the builder.
	WithSpecification(spec Specification[T]) SpecificationBuilder[T]

	// And combines the current specification with another specification using the logical AND operator.
	And(spec Specification[T]) SpecificationBuilder[T]

	// Or combines the current specification with another specification using the logical OR operator.
	Or(spec Specification[T]) SpecificationBuilder[T]

	// Not negates the current specification.
	Not() SpecificationBuilder[T]

	// Build constructs the final specification.
	Build() (Specification[T], error)
}