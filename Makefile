BASH_PATH:=$(shell which bash)
SHELL=$(BASH_PATH)
ROOT := $(shell realpath $(dir $(lastword $(MAKEFILE_LIST))))
APP := user-management
BUILD_PATH ?= ".build"
BUILD_DATE ?= $(shell TZ="Asia/Tehran" date +'%Y-%m-%dT%H:%M:%S%z')
COMPILER_VERSION ?= $(shell go version | cut -d' ' -f3)
DC_FILE="docker-compose.yml"
DC_RESOURCE_DIR=".compose"
CURRENT_TIMESTAMP := $(shell date +%s)

build:
	go build -v -race .

build-static:
	CGO_ENABLED=0 go build -v -o $(APP) -a -installsuffix  .

build-static-vendor: vendor
	CGO_ENABLED=0 go build -mod vendor -v -o $(APP) -installsuffix  .

build-static-vendor-linux: vendor
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod vendor -v -o $(APP) -installsuffix  .

bundle: build-static-vendor-linux
	mkdir -p ${BUILD_PATH}/assets/
	cp $(APP) ${BUILD_PATH}
	cp -r config.yaml ${BUILD_PATH}

install:
	cp $(APP) $(GOPATH)/bin

run:
	go run -race .

test:  vendor
	go test ./... -v

check-golint:
	go get -u golang.org/x/lint/golint

lint: check-golint
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R golint -set_exit_status R

check-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.2

lint-ci: check-golangci-lint vendor
	golangci-lint run -c .golangci.yml ./...

vendor:
	go mod vendor -v
