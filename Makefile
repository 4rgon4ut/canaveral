

APP_NAME = canaveral

COMMIT=$(shell git rev-parse --short HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
TAG=$(shell git describe --tags |cut -d- -f1)

LDFLAGS = -ldflags "-X main.gitTag=${TAG} -X main.gitCommit=${COMMIT} -X main.gitBranch=${BRANCH}"

.PHONY: help clean install-deps build lint

.DEFAULT_GOAL := help

help: ## Display this help screen.
	@echo "Makefile available targets:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  * \033[36m%-15s\033[0m %s\n", $$1, $$2}'


contract-tools:
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

go-mod:
	go mod download
	go mod tidy

dep: go-mod contract-tools
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0

install: dep  ## Build pgcenter executable.
	CGO_ENABLED=0 GOARCH=${GOARCH}
	go install cmd/canaveral.go

clean: ## Clean build directory.
	rm -rf ./artifacts/bin/
	rm -rf ./artifacts/abi/


lint:  ## Lint the source files
	golangci-lint run  --timeout 5m

test: go-mod
	go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...
	export PATH=$PATH:$(go env GOPATH)/bin
	ginkgo ./...