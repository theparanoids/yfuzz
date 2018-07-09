#
# Makefile to generate YFuzz docker images
#
# Copyright 2018 Oath, Inc.
# Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.
#

TARGET ?= $(shell basename `pwd`)
YFUZZ_BUILD_VERSION ?= $(shell git describe --tags --abbrev=0)_local

all: build

build:
	docker build -t ${TARGET} .

deps:

lint:
	docker run --rm -i hadolint/hadolint < Dockerfile

test:

clean:
