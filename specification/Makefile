# Variáveis
TIMEOUT = 30s
COVERAGE_PROFILE = coverage.out
TAGS = wireinject

# Alvo padrão
all: test

# Alvo para rodar testes
test:
	@echo "Running tests..."
	@go list -f '{{.Dir}}' ./... | grep -v '/test' | xargs -I {} sh -c 'if find {} -name "*_test.go" | grep -q .; then go test -tags $(TAGS) -failfast -timeout=$(TIMEOUT) -count=1 -cover -coverprofile=$(COVERAGE_PROFILE) -race {}; fi'

# Alvo para cobertura de código
coverage:
	@echo "Generating coverage report..."
	@go tool cover -html=$(COVERAGE_PROFILE)

# Limpeza
clean:
	@echo "Cleaning up..."
	@rm -f $(COVERAGE_PROFILE)
