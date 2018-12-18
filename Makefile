all: build

.PHONY: build

ifndef ($(GOPATH))
	GOPATH = $(HOME)/go
endif

PATH := $(GOPATH)/bin:$(PATH)
VERSION = $(shell git describe --tags --always --dirty)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
REVISION = $(shell git rev-parse HEAD)
REVSHORT = $(shell git rev-parse --short HEAD)
USER = $(shell whoami)
APP_NAME = gosal
PKGDIR_TMP = ${TMPDIR}golang

ifneq ($(OS), Windows_NT)
	CURRENT_PLATFORM = linux
	# If on macOS, set the shell to bash explicitly
	ifeq ($(shell uname), Darwin)
		SHELL := /bin/bash
		CURRENT_PLATFORM = darwin
	endif

	# To populate version metadata, we use unix tools to get certain data
	GOVERSION = $(shell go version | awk '{print $$3}')
	NOW	= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
else
	CURRENT_PLATFORM = windows

	# To populate version metadata, we use windows tools to get the certain data
	GOVERSION_CMD = "(go version).Split()[2]"
	GOVERSION = $(shell powershell $(GOVERSION_CMD))
	NOW	= $(shell powershell Get-Date -format s)
endif

BUILD_VERSION = "\
	-X github.com/airbnb/gosal/version.appName=${APP_NAME} \
	-X github.com/airbnb/gosal/version.version=${VERSION} \
	-X github.com/airbnb/gosal/version.branch=${BRANCH} \
	-X github.com/airbnb/gosal/version.buildUser=${USER} \
	-X github.com/airbnb/gosal/version.buildDate=${NOW} \
	-X github.com/airbnb/gosal/version.revision=${REVISION} \
	-X github.com/airbnb/gosal/version.goVersion=${GOVERSION}"

define HELP_TEXT

  Makefile commands

	make deps         - Install dependent programs and libraries
	make clean        - Delete all build artifacts

	make build        - Build the code

	make test         - Run the Go tests
	make lint         - Run the Go linters

endef

help:
	$(info $(HELP_TEXT))

deps:
	go get -u github.com/golang/dep/...
	go get -u github.com/golang/lint/golint
	dep ensure -vendor-only -v

clean:
	rm -rf build/
	rm -f *.zip
	rm -rf ${PKGDIR_TMP}_darwin
	rm -rf ${PKGDIR_TMP}_linux
	rm -rf ${PKGDIR_TMP}_windows

.pre-build:
	mkdir -p build/darwin
	mkdir -p build/linux
	mkdir -p build/windows


build: .pre-build
	GOOS=darwin go build -i -o build/darwin/${APP_NAME} -pkgdir ${PKGDIR_TMP}_darwin -ldflags ${BUILD_VERSION} ./cmd/gosal
	# GOOS=linux go build -i -o build/linux/${APP_NAME} -pkgdir ${PKGDIR_TMP}_linux -ldflags ${BUILD_VERSION} ./cmd/gosal
	GOOS=windows go build -i -o build/windows/${APP_NAME}.exe -pkgdir ${PKGDIR_TMP}_windows -ldflags ${BUILD_VERSION} ./cmd/gosal


test:
	go test -cover -race -v $(shell go list ./... | grep -v /vendor/)

lint:
	@if gofmt -l . | egrep -v ^vendor/ | grep .go; then \
	  echo "^- Repo contains improperly formatted go files; run gofmt -w *.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	go vet ./...
	# Bandaid until https://github.com/golang/lint/pull/325 is merged
	# TODO fix path to golint
	# golint -set_exit_status `go list ./... | grep -v /vendor/`
