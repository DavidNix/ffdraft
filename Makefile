default: help

.PHONY: help
help: ## Print this help message
	@echo "Available make commands:"; grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install
install: test ## go install
	go install ./...

test: ## go test
	@go test ./... -race --timeout=60s

.PHONY: projections
WEEK ?= 0
PPR ?= 0
projections: ## Use rscript to pull down projections
	@rscript R/projections.R --week $(WEEK) --ppr $(PPR)
