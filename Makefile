.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: build
build:
	go build

.PHONY: local_run
local_run: build
	./dsrhub
