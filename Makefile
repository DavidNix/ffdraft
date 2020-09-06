default: install

.PHONY: install
install: test
	go install ./...

test:
	@go test ./... -race --timeout=60s
