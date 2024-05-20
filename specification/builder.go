package specification

import (
	"errors"

	"github.com/mateusmacedo/vibranium/specification/contract"
)

// specificationBuilder is a struct that represents a builder for creating specifications.
type specificationBuilder[T contract.Candidate] struct {
	specifications []contract.Specification[T]
	err            error
}

// NewSpecificationBuilder creates a new instance of SpecificationBuilder for the given type T.
func NewSpecificationBuilder[T contract.Candidate]() contract.SpecificationBuilder[T] {
	return &specificationBuilder[T]{}
}

// addSpecification adds a specification to the specification builder.
// If there is an error in the builder or the given error is not nil,
// the error is set in the builder and the builder is returned.
// Otherwise, the specification is appended to the list of specifications
// in the builder and the builder is returned.
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

// WithSpecification sets the specification for the builder.
// It takes a spec of type contract.Specification[T] as a parameter.
// If the spec is nil, it returns an error indicating that the specification cannot be nil.
// Otherwise, it adds the specification to the builder and returns the updated builder.
func (b *specificationBuilder[T]) WithSpecification(spec contract.Specification[T]) contract.SpecificationBuilder[T] {
	if spec == nil {
		return b.addSpecification(nil, errors.New("specification cannot be nil"))
	}
	return b.addSpecification(spec, nil)
}

// And combines the current specification with the given specification using a logical AND operation.
// It returns a new SpecificationBuilder with the combined specification.
// If the given specification is nil, it returns an error indicating that the specification cannot be nil.
// If there is no previous specification to combine with, it returns an error indicating that there is no previous specification.
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

// Or combines the current specification with the provided specification using a logical OR operation.
// If the provided specification is nil, an error will be returned.
// If there are no existing specifications, the provided specification will be added as the first specification.
// Otherwise, the last specification will be combined with the provided specification using a logical OR operation.
// The combined specification will replace the last specification in the list.
// Returns the specification builder with the updated specifications.
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

// Not negates the last added specification in the builder.
// If there are no specifications to negate, it returns an error.
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

// Build returns the built specification based on the provided specifications.
// It returns an error if there was an error during the build process or if no specifications were provided.
// If only one specification was provided, it returns that specification directly.
// If multiple specifications were provided, it returns a new AndSpecification that combines all the provided specifications.
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
