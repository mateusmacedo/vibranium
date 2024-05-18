// file: pkg/and_specification_test.go
package specification

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/mateusmacedo/vibranium/specification/pkg/contract"
    mock_contract "github.com/mateusmacedo/vibranium/specification/tests/mocks/contract"
)

func TestAndSpecification_IsSatisfiedBy(t *testing.T) {
    tests := []struct {
        name      string
        candidate string
        specs     []func(candidate string) bool
        expected  bool
    }{
        {
            name:      "AllSpecificationsSatisfied",
            candidate: "candidate",
            specs: []func(candidate string) bool{
                func(candidate string) bool { return true },
                func(candidate string) bool { return true },
            },
            expected: true,
        },
        {
            name:      "OneSpecificationNotSatisfied",
            candidate: "candidate",
            specs: []func(candidate string) bool{
                func(candidate string) bool { return true },
                func(candidate string) bool { return false },
            },
            expected: false,
        },
        {
            name:      "NoSpecifications",
            candidate: "candidate",
            specs:     []func(candidate string) bool{},
            expected:  true,
        },
        {
            name:      "EmptyCandidate",
            candidate: "",
            specs: []func(candidate string) bool{
                func(candidate string) bool { return true },
                func(candidate string) bool { return true },
            },
            expected: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var specs []contract.Specification[string]
            for _, specFunc := range tt.specs {
                mockSpec := new(mock_contract.MockSpecification[string])
                mockSpec.On("IsSatisfiedBy", tt.candidate).Return(specFunc(tt.candidate))
                specs = append(specs, mockSpec)
            }

            andSpec := NewAndSpecification[string](specs...)

            // Act
            result := andSpec.IsSatisfiedBy(tt.candidate)

            // Assert
            assert.Equal(t, tt.expected, result)
            for _, spec := range specs {
                spec.(*mock_contract.MockSpecification[string]).AssertExpectations(t)
            }
        })
    }
}
