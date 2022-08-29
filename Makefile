

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

dep:
	go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...
	go get -u github.com/ethereum/go-ethereum
	cd $GOPATH/src/github.com/ethereum/go-ethereum/
	make
	make devtools
	sudo snap install solc --edge

go-mod:
	go mod download

build: dep go-mod ## Build pgcenter executable.
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ./${APP_NAME} ./cmd/main

clean: ## Clean build directory.
	rm -rf ./artifacts/bin/
	rm -rf ./artifacts/abi/


lint: dep go-mod ## Lint the source files
	golangci-lint run  --timeout 5m

test: dep go-mod
	ginkgo ./...