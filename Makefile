src-dirs = ./pkg/... ./init/... ./plugins/...

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

lint:
ifndef GOLANGCILINT
	@GO111MODULE=off go get -u github.com/myitcv/gobin
	@gobin github.com/golangci/golangci-lint/cmd/golangci-lint@v1.24.0
endif
	@golangci-lint run -D errcheck -E golint ${src-dirs}

test: lint
	go test -race -covermode=atomic ${src-dirs}

docker_build:
	docker-compose build

docker_run:
	docker-compose up --remove-orphans --force-recreate
