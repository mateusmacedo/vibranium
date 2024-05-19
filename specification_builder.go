package specification

import (
	"errors"

	"github.com/mateusmacedo/vibranium/specification/contract"
)

type specificationBuilder[T contract.Candidate] struct {
	specifications []contract.Specification[T]
	err            error
}

func NewSpecificationBuilder[T contract.Candidate]() contract.SpecificationBuilder[T] {
	return &specificationBuilder[T]{}
}

func (b *specificationBuilder[T]) addSpecification(spec contract.Specification[T], err error) contract.SpecificationBuilder[T] {
	if b.err != nil {
		return b
	}
	if err != nil {
		b.err = err
		return b
	}
	b.specifications = append(b.specifications, spec)
	return b
}

func (b *specificationBuilder[T]) WithSpecification(spec contract.Specification[T]) contract.SpecificationBuilder[T] {
	if spec == nil {
		return b.addSpecification(nil, errors.New("specification cannot be nil"))
	}
	return b.addSpecification(spec, nil)
}

func (b *specificationBuilder[T]) And(spec contract.Specification[T]) contract.SpecificationBuilder[T] {
	if spec == nil {
		return b.addSpecification(nil, errors.New("specification cannot be nil"))
	}
	if len(b.specifications) == 0 {
		return b.addSpecification(nil, errors.New("no previous specification to combine with"))
	}
	lastSpecIndex := len(b.specifications) - 1
	lastSpec := b.specifications[lastSpecIndex]
	combinedSpec := NewAndSpecification[T](lastSpec, spec)
	b.specifications = b.specifications[:lastSpecIndex]
	return b.addSpecification(combinedSpec, nil)
}

func (b *specificationBuilder[T]) Or(spec contract.Specification[T]) contract.SpecificationBuilder[T] {
	if spec == nil {
		return b.addSpecification(nil, errors.New("specification cannot be nil"))
	}
	if len(b.specifications) == 0 {
		return b.addSpecification(NewOrSpecification[T](spec), nil)
	}
	lastSpec := b.specifications[len(b.specifications)-1]
	combinedSpec := NewOrSpecification[T](lastSpec, spec)
	b.specifications = b.specifications[:len(b.specifications)-1]
	return b.addSpecification(combinedSpec, nil)
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
