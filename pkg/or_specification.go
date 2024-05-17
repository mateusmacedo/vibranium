package specification

import "github.com/mateusmacedo/vibranium/specification/pkg/contract"

type OrSpecification[T contract.Candidate] struct {
	specs []contract.Specification[T]
}

func NewOrSpecification[T contract.Candidate](specs ...contract.Specification[T]) *OrSpecification[T] {
	return &OrSpecification[T]{specs: specs}
}

func (s *OrSpecification[T]) IsSatisfiedBy(candidate T) bool {
	for _, spec := range s.specs {
		if spec.IsSatisfiedBy(candidate) {
			return true
		}
	}
	return false
}
