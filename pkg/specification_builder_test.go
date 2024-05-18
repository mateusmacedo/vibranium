package specification

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/vibranium/specification/pkg/contract"
)

type specificationBuilderTest struct {
	name      string
	candidate string
	expected  bool
	finalSpec func() (contract.Specification[string], error)
}

func TestSpecificationBuilderIntegration(t *testing.T) {
	tests := []specificationBuilderTest{
		{
			name:      "AllSpecificationsSatisfied",
			candidate: "candidate",
			expected:  true,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					WithSpecification(createMockSpec(true)).
					And(createMockSpec(true)).
					Or(createMockSpec(true)).
					Or(createMockSpec(true)).
					Not().
					Build()
			},
		},
		{
			name:      "OneAndSpecificationNotSatisfied",
			candidate: "candidate",
			expected:  false,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					WithSpecification(createMockSpec(true)).
					And(createMockSpec(false)).
					Or(createMockSpec(true)).
					Or(createMockSpec(true)).
					Not().
					Build()
			},
		},
		{
			name:      "OneOrSpecificationSatisfied",
			candidate: "candidate",
			expected:  true,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					WithSpecification(createMockSpec(true)).
					And(createMockSpec(true)).
					Or(createMockSpec(false)).
					Or(createMockSpec(true)).
					Not().
					Build()
			},
		},
		{
			name:      "AllOrSpecificationsNotSatisfied",
			candidate: "candidate",
			expected:  false,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					WithSpecification(createMockSpec(true)).
					And(createMockSpec(true)).
					Or(createMockSpec(false)).
					Or(createMockSpec(false)).
					Not().
					Build()
			},
		},
		{
			name:      "NotSpecificationSatisfied",
			candidate: "candidate",
			expected:  false,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					WithSpecification(createMockSpec(true)).
					And(createMockSpec(true)).
					Or(createMockSpec(true)).
					Or(createMockSpec(true)).
					Not().
					Build()
			},
		},
		{
			name:      "NoSpecifications",
			candidate: "candidate",
			expected:  false,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					Build()
			},
		},
		{
			name:      "EmptyCandidate",
			candidate: "",
			expected:  true,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					WithSpecification(createMockSpec(true)).
					And(createMockSpec(true)).
					Or(createMockSpec(true)).
					Or(createMockSpec(true)).
					Not().
					Build()
			},
		},
		// Novos cenários
		{
			name:      "AllAndSpecificationsNotSatisfied",
			candidate: "candidate",
			expected:  false,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					WithSpecification(createMockSpec(false)).
					And(createMockSpec(false)).
					Or(createMockSpec(true)).
					Or(createMockSpec(true)).
					Not().
					Build()
			},
		},
		{
			name:      "AllAndSpecificationsSatisfiedAndNotSpecSatisfied",
			candidate: "candidate",
			expected:  false,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					WithSpecification(createMockSpec(true)).
					And(createMockSpec(true)).
					Or(createMockSpec(false)).
					Or(createMockSpec(false)).
					Not().
					Build()
			},
		},
		{
			name:      "NoAndSpecificationsAndOrSpecSatisfied",
			candidate: "candidate",
			expected:  true,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					Or(createMockSpec(true)).
					Or(createMockSpec(false)).
					Not().
					Build()
			},
		},
		{
			name:      "NoAndSpecificationsAndOrSpecNotSatisfied",
			candidate: "candidate",
			expected:  false,
			finalSpec: func() (contract.Specification[string], error) {
				return NewSpecificationBuilder[string]().
					Or(createMockSpec(false)).
					Or(createMockSpec(false)).
					Not().
					Build()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec, err := tt.finalSpec()
			if err != nil {
				assert.Error(t, err)
			} else {
				result := spec.IsSatisfiedBy(tt.candidate)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}