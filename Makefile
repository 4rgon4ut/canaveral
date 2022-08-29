

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
ifeq (, $(shell which stringer))
	@echo "Installing stringer..."
	@go install golang.org/x/tools/cmd/stringer@latest
else
	@echo "stringer already installed; skipping..."
endif

ifeq (, $(shell which go-bindata))
	@echo "Installing go-bindata..."
	@go install github.com/kevinburke/go-bindata/go-bindata@latest
else
	@echo "go-bindata already installed; skipping..."
endif

ifeq (, $(shell which gencodec))
	@echo "Installing gencodec..."
	@go install github.com/fjl/gencodec@latest
else
	@echo "gencodec already installed; skipping..."
endif

ifeq (, $(shell which protoc-gen-go))
	@echo "Installing protoc-gen-go..."
	@go install github.com/fjl/gencodeclatest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
else
	@echo "protoc-gen-go already installed; skipping..."
endif

ifeq (, $(shell which protoc-gen-go-grpc))
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
else
	@echo "protoc-gen-go-grpc already installed; skipping..."
endif

ifeq (, $(shell which solhint))
	@echo "Installing solhint..."
	@npm install -g solhint
else
	@echo "solhint already installed; skipping..."
endif

ifeq (, $(shell which solcjs))
	@echo "Installing solcjs..."
	@npm install -g solc@0.5.11
else
	@echo "solcjs already installed; skipping..."
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

build: dep  ## Build pgcenter executable.
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ./${APP_NAME} ./cmd/main

clean: ## Clean build directory.
	rm -rf ./artifacts/bin/
	rm -rf ./artifacts/abi/


lint: dep  ## Lint the source files
	golangci-lint run  --timeout 5m

test: dep
	ginkgo ./...