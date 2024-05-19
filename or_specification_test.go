// file: pkg/or_specification_test.go
package specification

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mateusmacedo/vibranium/specification/contract"
	"github.com/mateusmacedo/vibranium/specification/mocks"
)

func TestOrSpecification_IsSatisfiedBy(t *testing.T) {
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
            name:      "OneSpecificationSatisfied",
            candidate: "candidate",
            specs: []func(candidate string) bool{
                func(candidate string) bool { return false },
                func(candidate string) bool { return true },
            },
            expected: true,
        },
        {
            name:      "NoSpecificationsSatisfied",
            candidate: "candidate",
            specs: []func(candidate string) bool{
                func(candidate string) bool { return false },
                func(candidate string) bool { return false },
            },
            expected: false,
        },
        {
            name:      "NoSpecifications",
            candidate: "candidate",
            specs:     []func(candidate string) bool{},
            expected:  false,
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
            var specmocks []*mocks.MockSpecification[string]

            for _, specFunc := range tt.specs {
                mockSpec := new(mocks.MockSpecification[string])
                mockSpec.On("IsSatisfiedBy", tt.candidate).Return(specFunc(tt.candidate)).Once()
                specs = append(specs, mockSpec)
                specmocks = append(specmocks, mockSpec)
            }

            orSpec := NewOrSpecification[string](specs...)

            // Act
            result := orSpec.IsSatisfiedBy(tt.candidate)

            // Assert
            assert.Equal(t, tt.expected, result)

            // Verificar expectativas
            for i, mock := range specmocks {
                if tt.expected && i > 0 && tt.specs[i-1](tt.candidate) {
                    // Se o resultado esperado for verdadeiro e a especificação anterior retornou verdadeiro,
                    // verificamos que esta especificação não foi chamada
                    mock.AssertNotCalled(t, "IsSatisfiedBy", tt.candidate)
                } else {
                    // Caso contrário, verificamos que a especificação foi chamada
                    mock.AssertExpectations(t)
                }
            }
        })
    }
}
