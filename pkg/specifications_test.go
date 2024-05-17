package specification

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/vibranium/specification/pkg/contract"
	mock_contract "github.com/mateusmacedo/vibranium/specification/tests/mocks/contract"
)

type specTest struct {
	name      string
	candidate string
	expected  bool
	finalSpec contract.Specification[string]
}

func TestSpecificationIntegration(t *testing.T) {
	tests := []specTest{
		{
			name:      "AllSpecificationsSatisfied",
			candidate: "candidate",
			expected:  true,
			finalSpec: NewAndSpecification[string](
				createMockSpec(true),
				createMockSpec(true),
				NewOrSpecification[string](
					createMockSpec(true),
					createMockSpec(true),
				),
				NewNotSpecification[string](createMockSpec(false)),
			),
		},
		{
			name:      "OneAndSpecificationNotSatisfied",
			candidate: "candidate",
			expected:  false,
			finalSpec: NewAndSpecification[string](
				createMockSpec(true),
				createMockSpec(false),
				NewOrSpecification[string](
					createMockSpec(true),
					createMockSpec(true),
				),
				NewNotSpecification[string](createMockSpec(false)),
			),
		},
		{
			name:      "OneOrSpecificationSatisfied",
			candidate: "candidate",
			expected:  true,
			finalSpec: NewAndSpecification[string](
				createMockSpec(true),
				createMockSpec(true),
				NewOrSpecification[string](
					createMockSpec(false),
					createMockSpec(true),
				),
				NewNotSpecification[string](createMockSpec(false)),
			),
		},
		{
			name:      "AllOrSpecificationsNotSatisfied",
			candidate: "candidate",
			expected:  false,
			finalSpec: NewAndSpecification[string](
				createMockSpec(true),
				createMockSpec(true),
				NewOrSpecification[string](
					createMockSpec(false),
					createMockSpec(false),
				),
				NewNotSpecification[string](createMockSpec(false)),
			),
		},
		{
			name:      "NotSpecificationSatisfied",
			candidate: "candidate",
			expected:  false,
			finalSpec: NewAndSpecification[string](
				createMockSpec(true),
				createMockSpec(true),
				NewOrSpecification[string](
					createMockSpec(true),
					createMockSpec(true),
				),
				NewNotSpecification[string](createMockSpec(true)),
			),
		},
		{
			name:      "NoSpecifications",
			candidate: "candidate",
			expected:  true,
			finalSpec: NewAndSpecification[string](
				NewOrSpecification[string](
					NewNotSpecification[string](createMockSpec(false)),
				),
			),
		},
		{
			name:      "EmptyCandidate",
			candidate: "",
			expected:  true,
			finalSpec: NewAndSpecification[string](
				createMockSpec(true),
				createMockSpec(true),
				NewOrSpecification[string](
					createMockSpec(true),
					createMockSpec(true),
				),
				NewNotSpecification[string](createMockSpec(false)),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.finalSpec.IsSatisfiedBy(tt.candidate)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func createMockSpec(result bool) contract.Specification[string] {
	mockSpec := new(mock_contract.MockSpecification[string])
	mockSpec.On("IsSatisfiedBy", mock.Anything).Return(result)
	return mockSpec
}