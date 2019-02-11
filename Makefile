# Not required right now 
# GOOS=$(shell go env GOOS)
# GOARCH=$(shell go env GOARCH)

BINARY_BASE_NAME=binary

.PHONY: build
build: 
	go build -o ./target/$(BINARY_BASE_NAME)

.PHONY: build-all
build-all: 
	GOOS=linux GOARCH=amd64 go build -o ./target/$(BINARY_BASE_NAME)-linux
	GOOS=darwin GOARCH=amd64 go build -o ./target/$(BINARY_BASE_NAME)-mac
	GOOS=linux GOARCH=amd64 go build -o ./target/$(BINARY_BASE_NAME)-windows
