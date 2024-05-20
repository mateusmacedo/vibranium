# Variables
TIMEOUT = 30s
COVERAGE_PROFILE = coverage.out
TAGS = wireinject
EXCLUDE_PATTERN = "mocks|contract"
TMP_DIR := $(shell mktemp -d)

# Default target
all: test

# Target for running tests
test:
	@echo "Running tests..."
	@trap 'rm -rf $(TMP_DIR)' EXIT  # Ensure cleanup happens on script exit
	@echo "mode: atomic" > $(COVERAGE_PROFILE)  # Initialize the coverage profile
	@go list ./... | grep -vE $(EXCLUDE_PATTERN) | while read pkg; do \
		mkdir -p $(TMP_DIR)/$$pkg; \
		go test -tags $(TAGS) -failfast -timeout=$(TIMEOUT) -count=1 -covermode=atomic -coverprofile=$(TMP_DIR)/$$pkg/coverage.out -race $$pkg; \
		if [ -f $(TMP_DIR)/$$pkg/coverage.out ]; then \
			cat $(TMP_DIR)/$$pkg/coverage.out | grep -v "mode: atomic" >> $(COVERAGE_PROFILE); \
			rm -rf $(TMP_DIR)/$$pkg; \
		fi; \
	done

# Target for code coverage
coverage: test
	@echo "Generating coverage report..."
	@go tool cover -html=$(COVERAGE_PROFILE)

# Cleanup
clean:
	@echo "Cleaning up..."
	@rm -f $(COVERAGE_PROFILE)
	@rm -rf $(TMP_DIR)