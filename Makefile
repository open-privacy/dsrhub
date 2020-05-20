GOLANGCILINT := $(shell command -v golangci-lint 2> /dev/null)

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

lint:
ifndef GOLANGCILINT
	@GO111MODULE=off go get -u github.com/myitcv/gobin
	@gobin github.com/golangci/golangci-lint/cmd/golangci-lint@v1.24.0
endif
	@golangci-lint run -D errcheck .

test: lint
	go test -race -covermode=atomic ./pkg/...

build:
	go build

local_run: build
	./dsrhub
