
BIN ?= dist/srv

build: ## Builds the binary
	go build -o $(BIN)
.PHONY: build

test: ## Run unit tests
	go test -v ./...
.PHONY: test
