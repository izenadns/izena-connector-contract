.DEFAULT_GOAL := help

.PHONY: help fmt fmt-check lint security module-check test ci

help: ## Show the contract development commands.
	@awk 'BEGIN {FS = ":.*## "; print "Contract development commands:"} /^[a-zA-Z0-9_-]+:.*## / {printf "  %-26s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

fmt: ## Format Go source.
	@go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15.0 fmt

fmt-check: ## Check Go formatting.
	@go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15.0 fmt --diff

lint: ## Run configured Go linters.
	@go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15.0 run

security: ## Scan Go dependencies.
	@go run github.com/google/osv-scanner/v2/cmd/osv-scanner@v2.7.1 scan source \
		--config osv-scanner.toml --no-call-analysis=go --lockfile go.mod

module-check: ## Check standalone module dependencies.
	@GOWORK=off go mod tidy -diff

test: ## Run unit tests with the race detector.
	@go test -race ./...

ci: fmt-check lint security module-check test ## Run the contract CI checks.
