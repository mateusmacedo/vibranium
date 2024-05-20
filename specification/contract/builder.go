package contract

// SpecificationBuilder is an interface for building specifications.
type SpecificationBuilder[T Candidate] interface {
	// WithSpecification adds a specification to the builder.
	//
	// Parameters:
	// - spec: The specification to be added.
	//
	// Returns:
	// - SpecificationBuilder[T]: The builder itself to allow for method chaining.
	WithSpecification(spec Specification[T]) SpecificationBuilder[T]

	// And combines the current specification with another specification using the logical AND operator.
	//
	// Parameters:
	// - spec: The specification to be combined using AND.
	//
	// Returns:
	// - SpecificationBuilder[T]: The builder itself to allow for method chaining.
	And(spec Specification[T]) SpecificationBuilder[T]

	// Or combines the current specification with another specification using the logical OR operator.
	//
	// Parameters:
	// - spec: The specification to be combined using OR.
	//
	// Returns:
	// - SpecificationBuilder[T]: The builder itself to allow for method chaining.
	Or(spec Specification[T]) SpecificationBuilder[T]

	// Not negates the current specification.
	//
	// Returns:
	// - SpecificationBuilder[T]: The builder itself to allow for method chaining.
	Not() SpecificationBuilder[T]

	// Build constructs the final specification.
	//
	// Returns:
	// - Specification[T]: The final specification.
	// - error: An error if the specification cannot be built.
	Build() (Specification[T], error)
}
