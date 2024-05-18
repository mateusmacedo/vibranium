// file: pkg/not_specification_test.go
package specification

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mock_contract "github.com/mateusmacedo/vibranium/specification/tests/mocks/contract"
)

func TestNotSpecification_IsSatisfiedBy(t *testing.T) {
    tests := []struct {
        name      string
        candidate string
        specFunc  func(candidate string) bool
        expected  bool
    }{
        {
            name:      "UnderlyingSpecSatisfied",
            candidate: "candidate",
            specFunc:  func(candidate string) bool { return true },
            expected:  false,
        },
        {
            name:      "UnderlyingSpecNotSatisfied",
            candidate: "candidate",
            specFunc:  func(candidate string) bool { return false },
            expected:  true,
        },
        {
            name:      "EmptyCandidate",
            candidate: "",
            specFunc:  func(candidate string) bool { return false },
            expected:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockSpec := new(mock_contract.MockSpecification[string])
            mockSpec.On("IsSatisfiedBy", tt.candidate).Return(tt.specFunc(tt.candidate))

            notSpec := NewNotSpecification[string](mockSpec)

            // Act
            result := notSpec.IsSatisfiedBy(tt.candidate)

            // Assert
            assert.Equal(t, tt.expected, result)
            mockSpec.AssertExpectations(t)
        })
    }
}
