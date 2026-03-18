.PHONY: build clean install test fmt lint help

BINARY_NAME=zohodesk-cli
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION)-X main.buildTime=$(BUILD_TIME)"

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build binary
	go build $(LDFLAGS) -o $(BINARY_NAME) .

install: build ## Install to /usr/local/bin
	sudo mv $(BINARY_NAME) /usr/local/bin/

test: ## Run tests
	go test -v -race ./...

fmt: ## Format code
	go fmt ./...

lint: ## Run linter
	golangci-lint run ./...

clean: ## Clean build artifacts
	rm -f $(BINARY_NAME)
	rm -rf dist/

run: build ## Run locally
	./$(BINARY_NAME)

snapshot: ## Create snapshot release (local)
	goreleaser release --snapshot --clean

build-all: ## Build for all platforms
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64.
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe .

docker-build: ## Build Docker image
	docker build -t $(BINARY_NAME):latest .

docker-run: docker-build ## Run Docker container
	docker run --rm -it $(BINARY_NAME):latest version

deps: ## Download dependencies
	go mod download
	go mod tidy