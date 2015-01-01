all: clean test

setup:
	go get github.com/golang/lint/golint

test:
	go test $(TESTFLAGS) ./resize ./webp ./server

clean:
	go clean

lint:
	golint ./...

vet:
	go vet ./...

.PHONY: setup test clean lint vet