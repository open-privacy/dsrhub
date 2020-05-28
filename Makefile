.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
