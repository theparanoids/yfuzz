#
# Makefile to generate YFuzz server binary
#
# Copyright 2018 Oath, Inc.
# Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.
#
BINARY_NAME = yfuzz-server
YFUZZ_BUILD_VERSION ?= $(shell git describe --tags `git rev-parse HEAD`)_local

# Run go build or go install with the appropriate flags
define _build
	@echo Building YFuzz Version ${YFUZZ_BUILD_VERSION}
	go $(1) -ldflags "-s -w -X github.com/yahoo/yfuzz/pkg/version.Version=${YFUZZ_BUILD_VERSION} \
													-X github.com/yahoo/yfuzz/pkg/version.Timestamp=$(shell date +'%Y/%m/%d_%H:%M:%S')"
endef

all: deps lint test build

clean:
	go clean
	rm -f yfuzz-server

deps:
	glide install
ifndef TRAVIS
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
endif

lint:
ifdef TRAVIS
	@echo Skipping lint in Travis, linting is done by https://golangci.com/
else
	golangci-lint run -E golint -E interfacer -E unconvert -E goconst -E goimports ./...
endif

test:
	go test -v ./...

build:
	$(call _build,build -o ${BINARY_NAME})

install:
	$(call _build,"install")

.PHONY: clean deps lint test build install
