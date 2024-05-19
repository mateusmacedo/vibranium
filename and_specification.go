package specification

import "github.com/mateusmacedo/vibranium/specification/contract"

type AndSpecification[T contract.Candidate] struct {
    specs []contract.Specification[T]
}

func NewAndSpecification[T contract.Candidate](specs ...contract.Specification[T]) *AndSpecification[T] {
    return &AndSpecification[T]{specs: specs}
}

func (s *AndSpecification[T]) IsSatisfiedBy(candidate T) bool {
    for _, spec := range s.specs {
        if !spec.IsSatisfiedBy(candidate) {
            return false
        }
    }
    return true
}