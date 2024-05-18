package specification

import (
	"errors"

	"github.com/mateusmacedo/vibranium/specification/pkg/contract"
)

type specificationBuilder[T any] struct {
	specifications []contract.Specification[T]
	err            error
}

func NewSpecificationBuilder[T any]() contract.SpecificationBuilder[T] {
	return &specificationBuilder[T]{}
}

func (b *specificationBuilder[T]) WithSpecification(spec contract.Specification[T]) contract.SpecificationBuilder[T] {
	if b.err != nil {
		return b
	}
	if spec == nil {
		b.err = errors.New("specification cannot be nil")
		return b
	}
	b.specifications = append(b.specifications, spec)
	return b
}

func (b *specificationBuilder[T]) And(spec contract.Specification[T]) contract.SpecificationBuilder[T] {
	if b.err != nil {
		return b
	}
	if spec == nil {
		b.err = errors.New("specification cannot be nil")
		return b
	}
	b.specifications = append(b.specifications, NewAndSpecification[T](spec))
	return b
}

func (b *specificationBuilder[T]) Or(spec contract.Specification[T]) contract.SpecificationBuilder[T] {
	if b.err != nil {
		return b
	}
	if spec == nil {
		b.err = errors.New("specification cannot be nil")
		return b
	}
	b.specifications = append(b.specifications, NewOrSpecification[T](spec))
	return b
}

func (b *specificationBuilder[T]) Not() contract.SpecificationBuilder[T] {
	if b.err != nil {
		return b
	}
	if len(b.specifications) == 0 {
		b.err = errors.New("no specification to negate")
		return b
	}
	lastSpec := b.specifications[len(b.specifications)-1]
	b.specifications = b.specifications[:len(b.specifications)-1]
	b.specifications = append(b.specifications, NewNotSpecification[T](lastSpec))
	return b
}

func (b *specificationBuilder[T]) Build() (contract.Specification[T], error) {
	if b.err != nil {
		return nil, b.err
	}
	if len(b.specifications) == 0 {
		return nil, errors.New("no specifications provided")
	}
	if len(b.specifications) == 1 {
		return b.specifications[0], nil
	}
	return NewAndSpecification[T](b.specifications...), nil
}