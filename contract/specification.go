package contract

type Candidate any

// Specification is an interface that defines a contract for a specification.
// A specification is used to determine if a candidate object satisfies certain criteria.
type Specification[T Candidate] interface {
	IsSatisfiedBy(candidate T) bool
}