src-dirs = ./pkg/... ./init/... ./plugins/...
CMD_GOLANGCILINT := $(shell command -v golangci-lint 2> /dev/null)
CMD_SWAGGER := $(shell command -v swagger 2> /dev/null)

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

lint: _install_cmd_golangci_lint _install_cmd_swagger
	golangci-lint run -D errcheck -E golint $(src-dirs)
	swagger validate $(PWD)/idl_dsrhub/dsrhub.swagger.json

test: lint
	go test -race -covermode=atomic $(src-dirs)

docker_build:
	docker-compose build

docker_run:
	docker-compose up --remove-orphans --force-recreate

docker_functional_test:
	docker-compose down -v
	docker-compose build
	docker-compose up --no-start
	for f in ./mocks/openmock_templates/*.yaml; do docker cp $$f dsrhub_openmock:/data/templates/; done
	for f in ./templates/*.yaml;                do docker cp $$f dsrhub_utask:/app/templates/; done
	for f in ./functional_test/*.yaml;          do docker cp $$f dsrhub_functional_test:/var/local/; done
	docker-compose up -d
	docker-compose exec -w /var/local functional_test venom run --log info '*'

gen_proto:
	docker run --rm \
      -v $(PWD)/idl_dsrhub:/defs \
      -v $(PWD)/idl_dsrhub:/go/src/github.com/dsrhub/dsrhub/idl_dsrhub \
      -v $(PWD)/idl_dsrhub/dsrhub.swagger.json:/go/src/dsrhub.swagger.json \
      namely/protoc-all:1.29_2 -i . -d . -l go -o /go/src --with-gateway

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

