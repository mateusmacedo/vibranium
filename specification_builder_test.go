package specification

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mateusmacedo/vibranium/specification/contract"
	mock_contract "github.com/mateusmacedo/vibranium/specification/tests/mocks/contract"
)

func TestSpecificationBuilderCases(t *testing.T) {
	tests := []struct {
		name      string
		candidate string
		build     func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string]
		expected  bool
		expectErr bool
	}{
		{
			name:      "AllSpecificationsSatisfied",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(createMockSpecForBuilder(true)).
					WithSpecification(createMockSpecForBuilder(true))
			},
			expected: true,
		},
		{
			name:      "OneAndSpecificationNotSatisfied",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(createMockSpecForBuilder(true)).
					And(createMockSpecForBuilder(false))
			},
			expected: false,
		},
		{
			name:      "OneOrSpecificationSatisfied",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(createMockSpecForBuilder(false)).
					Or(createMockSpecForBuilder(true))
			},
			expected: true,
		},
		{
			name:      "AllOrSpecificationsNotSatisfied",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(createMockSpecForBuilder(false)).
					Or(createMockSpecForBuilder(false))
			},
			expected: false,
		},
		{
			name:      "NotSpecificationSatisfied",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(createMockSpecForBuilder(true)).
					Not()
			},
			expected: false,
		},
		{
			name:      "NoSpecifications",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder
			},
			expected: false,
			expectErr: true,
		},
		{
			name:      "EmptyCandidate",
			candidate: "",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(createMockSpecForBuilder(true)).
					WithSpecification(createMockSpecForBuilder(true))
			},
			expected: true,
		},
		{
			name:      "AllAndSpecificationsNotSatisfied",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(createMockSpecForBuilder(false)).
					And(createMockSpecForBuilder(false))
			},
			expected: false,
		},
		{
			name:      "AllAndSpecificationsSatisfiedAndNotSpecSatisfied",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(createMockSpecForBuilder(true)).
					And(createMockSpecForBuilder(true)).
					Not()
			},
			expected: false,
		},
		{
			name:      "NoAndSpecificationsAndOrSpecSatisfied",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.Or(createMockSpecForBuilder(true))
			},
			expected: true,
		},
		{
			name:      "NoAndSpecificationsAndOrSpecNotSatisfied",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.Or(createMockSpecForBuilder(false))
			},
			expected: false,
		},
		{
			name:      "WithNilSpecification",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(nil)
			},
			expectErr: true,
		},
		{
			name:      "AndWithNilSpecification",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(createMockSpecForBuilder(true)).And(nil)
			},
			expectErr: true,
		},
		{
			name:      "OrWithNilSpecification",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(createMockSpecForBuilder(true)).Or(nil)
			},
			expectErr: true,
		},
		{
			name:      "NotWithoutSpecification",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.Not()
			},
			expectErr: true,
		},
		{
			name:      "NotAfterError",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(nil).Not()
			},
			expectErr: true,
		},
		{
			name:      "AndAfterError",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(nil).And(createMockSpecForBuilder(true))
			},
			expectErr: true,
		},
		{
			name:      "OrAfterError",
			candidate: "candidate",
			build: func(builder contract.SpecificationBuilder[string]) contract.SpecificationBuilder[string] {
				return builder.WithSpecification(nil).Or(createMockSpecForBuilder(true))
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewSpecificationBuilder[string]()
			finalBuilder := tt.build(builder)
			finalSpec, err := finalBuilder.Build()

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, finalSpec)
			} else {
				assert.NoError(t, err)
				result := finalSpec.IsSatisfiedBy(tt.candidate)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func createMockSpecForBuilder(result bool) contract.Specification[string] {
	mockSpec := new(mock_contract.MockSpecification[string])
	mockSpec.On("IsSatisfiedBy", mock.Anything).Return(result)
	return mockSpec
}
