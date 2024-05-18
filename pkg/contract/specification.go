package contract

type Candidate any

type Specification[T Candidate] interface {
	IsSatisfiedBy(candidate T) bool
}