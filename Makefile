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
	@golangci-lint run -D errcheck -E golint ./pkg/...
	@golangci-lint run -D errcheck -E golint ./init/...
	@golangci-lint run -D errcheck -E golint ./plugins/...

test: lint
	go test -race -covermode=atomic ./pkg/...

build_plugins:
	for p in $$(ls ./plugins); do if [ -d ./plugins/$$p ]; then GO111MODULE=on go build --buildmode=plugin -o ./plugins/$$p.so ./plugins/$$p/*.go; fi; done
	for p in $$(ls ./init);    do if [ -d ./init/$$p ];    then GO111MODULE=on go build --buildmode=plugin -o ./init/$$p.so    ./init/$$p/*.go;    fi; done

build: build_plugins
