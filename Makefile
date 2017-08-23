all: build

deps:
	go get -u github.com/golang/dep/...
	go get -u github.com/golang/lint/golint
	dep ensure -v

build:
	mkdir -p build
	go build -i -o build/gosal
	GOOS=windows go build -o build/gosal.exe

test:
	go test -cover -race -v $(shell go list ./... | grep -v /vendor/)

.PHONY: build

lint:
	@if gofmt -l . | egrep -v ^vendor/ | grep .go; then \
	  echo "^- Repo contains improperly formatted go files; run gofmt -w *.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	for pkg in $$(go list ./... |grep -v /vendor/); do golint $$pkg; done
