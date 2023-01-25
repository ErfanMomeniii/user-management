BASH_PATH:=$(shell which bash)
SHELL=$(BASH_PATH)
ROOT := $(shell realpath $(dir $(lastword $(MAKEFILE_LIST))))
APP := onefootballTask
BUILD_PATH ?= ".build"
BUILD_DATE ?= $(shell TZ="Asia/Tehran" date +'%Y-%m-%dT%H:%M:%S%z')
COMPILER_VERSION ?= $(shell go version | cut -d' ' -f3)
DC_FILE="docker-compose.yml"
DC_RESOURCE_DIR=".compose"
CURRENT_TIMESTAMP := $(shell date +%s)

set-goproxy:
	go env -w GOPROXY=""
	go env -w GONOSUMDB=""

build: set-goproxy
	go build -v -race .

build-static: set-goproxy
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

run: set-goproxy
	go run -race .

check-gotestsum: set-goproxy
	which gotestsum || (go get -u gotest.tools/gotestsum)

test: check-gotestsum vendor
	gotestsum --junitfile-testcase-classname short --junitfile .report.xml -- -gcflags 'all=-N -l' -mod vendor ./...

coverage: vendor
	gotestsum -- -gcflags 'all=-N -l' -mod vendor -v -coverprofile=.testCoverage.txt ./...
	GOFLAGS=-mod=vendor go tool cover -func=.testCoverage.txt

coverage-report: coverage
	GOFLAGS=-mod=vendor go tool cover -html=.testCoverage.txt -o testCoverageReport.html
	gocover-cobertura < .testCoverage.txt > .cobertura.xml

check-golint: set-goproxy
	which golint || (go get -u golang.org/x/lint/golint)

lint: check-golint
	find $(ROOT) -type f -name "*.go" -not -path "$(ROOT)/vendor/*" | xargs -n 1 -I R golint -set_exit_status R

check-golangci-lint: set-goproxy
	which golangci-lint || (go get -u github.com/golangci/golangci-lint/cmd/golangci-lint)

lint-ci: check-golangci-lint vendor
	golangci-lint run -c .golangci.yml ./...

vendor:
	go mod vendor -v
