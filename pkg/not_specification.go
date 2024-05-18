package specification

import "github.com/mateusmacedo/vibranium/specification/pkg/contract"

type NotSpecification[T contract.Candidate] struct {
    spec contract.Specification[T]
}

func NewNotSpecification[T contract.Candidate](spec contract.Specification[T]) *NotSpecification[T] {
    return &NotSpecification[T]{spec: spec}
}

func (s *NotSpecification[T]) IsSatisfiedBy(candidate T) bool {
    return !s.spec.IsSatisfiedBy(candidate)
}