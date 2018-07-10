#
# Makefile to generate YFuzz docker images
#
# Copyright 2018 Oath, Inc.
# Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.
#

TARGET ?= $(shell basename `pwd`)
YFUZZ_BUILD_VERSION ?= $(shell git describe --tags --abbrev=0)_local
GIT_SHA := $(shell git rev-parse --short HEAD)
IMAGE_NAME := yfuzz/${TARGET}:build-${GIT_SHA}

all: lint build push

deps:

lint:
	docker run --rm -i hadolint/hadolint < Dockerfile

build:
	docker build -t ${IMAGE_NAME} .

test:

push:
ifdef TRAVIS
	@echo "Pushing image ${IMAGE_NAME} to docker hub."
	@echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
	docker push ${IMAGE_NAME}
endif

clean:

.PHONY: build deps lint test push clean