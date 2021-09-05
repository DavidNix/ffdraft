default: install

.PHONY: install
install: test
	go install ./...

test:
	@go test ./... -race --timeout=60s

.PHONY: projections
WEEK ?= 0
PPR ?= 0
projections:
	@rscript R/projections.R --week $(WEEK) --ppr $(PPR)
