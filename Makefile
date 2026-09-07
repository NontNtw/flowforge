.DEFAULT_GOAL := help

.PHONY: help
help: ## Show available commands
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: check
check: ## Run repository validation checks
	@test -f README.md
	@test -f CONTRIBUTING.md
	@test -f .gitignore
	@test -f .github/pull_request_template.md
	@test -d apps/api
	@test -d apps/worker
	@test -d internal
	@test -d platform
	@test -d migrations
	@test -d sql/queries
	@test -d sql/schema
	@echo "Repository validation passed."

.PHONY: fmt
fmt: ## Format Go source code
	@if find . -name '*.go' -type f | grep -q .; then \
		gofmt -w $$(find . -name '*.go' -type f); \
	else \
		echo "No Go files to format yet."; \
	fi

.PHONY: test
test: ## Run Go tests
	@if [ -f go.mod ]; then \
		go test ./...; \
	else \
		echo "Go module has not been initialized yet."; \
	fi

.PHONY: vet
vet: ## Run go vet
	@if [ -f go.mod ]; then \
		go vet ./...; \
	else \
		echo "Go module has not been initialized yet."; \
	fi

.PHONY: ci
ci: check fmt vet test ## Run local CI checks
