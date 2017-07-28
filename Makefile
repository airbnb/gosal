all: build

deps:
	go get -u github.com/golang/dep/...
	dep ensure -v

build:
	mkdir -p build
	go build -i -o build/gosal
	GOOS=windows go build -i -o build/gosal.exe

test:
	go test -cover -race -v $(shell go list ./... | grep -v /vendor/)

.PHONY: build 
