

APP_NAME = canaveral

COMMIT=$(shell git rev-parse --short HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
TAG=$(shell git describe --tags |cut -d- -f1)

LDFLAGS = -ldflags "-X main.gitTag=${TAG} -X main.gitCommit=${COMMIT} -X main.gitBranch=${BRANCH}"

.PHONY: help clean deps build lint install

.DEFAULT_GOAL := help

help: ## Display this help screen.
	@echo "Makefile available targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  * \033[36m%-15s\033[0m %s\n", $$1, $$2}'


contract-tools: ## Install ethereum specific tools (solc, solhing, abigen, ...)
ifeq (, $(shell which go-bindata))
	@echo "Installing go-bindata..."
	@go install github.com/kevinburke/go-bindata/go-bindata@latest
else
	@echo "go-bindata already installed; skipping..."
endif
ifeq (, $(shell which solhint))
	@echo "Installing solhint..."
	@npm install -g solhint
else
	@echo "solhint already installed; skipping..."
endif
ifeq (, $(shell which solc))
	@echo "Installing solc..."
	@snap install solc
else
	@echo "solc already installed; skipping..."
endif
 ifeq (, $(shell which abigen))
	@echo "Installing abigen..."
	@add-apt-repository -y ppa:ethereum/ethereum
	@apt-get install -y ethereum
else
	@echo "abigen already installed; skipping..."
endif

deps: contract-tools ## Install all necessary dependencie
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	go mod download
	go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...
	export PATH=$PATH:$(go env GOPATH)/bin
	go mod tidy

lint: deps ## Runs linting tool
	golangci-lint run  --timeout 5m

test: deps ## Runs all tests
	ginkgo ./...
	go mod tidy

install: deps ## Install canaveral cli to go binaries dir
	CGO_ENABLED=0 GOARCH=${GOARCH}
	go install cmd/canaveral.go
	export GOPATH="$HOME/go"
	export PATH="$PATH:$GOPATH/bin"

clean: ## Clean generated artifacts
	rm -rf ./artifacts/bin/
	rm -rf ./artifacts/abi/
