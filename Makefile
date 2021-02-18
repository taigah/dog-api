.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o dogapi bin/cli/cli.go