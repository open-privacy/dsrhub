src-dirs = ./pkg/... ./init/... ./plugins/...
CMD_GOLANGCILINT := $(shell command -v golangci-lint 2> /dev/null)
CMD_SWAGGER := $(shell command -v swagger 2> /dev/null)

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

lint: _install_cmd_golangci_lint _install_cmd_swagger
	golangci-lint run -D errcheck -E golint $(src-dirs)
	swagger validate $(PWD)/idl_dsrhub/swagger_opendsr.yaml

test: lint
	go test -race -covermode=atomic $(src-dirs)

docker_build:
	docker-compose build

docker_run:
	docker-compose up --remove-orphans --force-recreate --scale utask=2

gen_proto:
	docker run --rm \
      -v $(PWD)/idl_dsrhub:/defs \
      -v $(PWD)/idl_dsrhub:/go/src/github.com/dsrhub/dsrhub/idl_dsrhub \
      namely/protoc-all:1.11 -i . -d . -l go -o /go/src

_install_cmd_golangci_lint:
ifndef CMD_GOLANGCILINT
	GO111MODULE=off go get -u github.com/myitcv/gobin
	gobin github.com/golangci/golangci-lint/cmd/golangci-lint@v1.24.0
endif

_install_cmd_swagger:
ifndef CMD_SWAGGER
	GO111MODULE=off go get -u github.com/myitcv/gobin
	gobin github.com/go-swagger/go-swagger/cmd/swagger@v0.24.0
endif

