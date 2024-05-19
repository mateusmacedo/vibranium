package contract

type SpecificationBuilder[T Candidate] interface {
	WithSpecification(spec Specification[T]) SpecificationBuilder[T]
	And(spec Specification[T]) SpecificationBuilder[T]
	Or(spec Specification[T]) SpecificationBuilder[T]
	Not() SpecificationBuilder[T]
	Build() (Specification[T], error)
}