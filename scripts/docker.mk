#
# Makefile to generate yFuzz docker images
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
	@echo "deps: nothing to do for image ${IMAGE_NAME}."

lint:
	@echo Running linters for ${IMAGE_NAME}.
	docker run --rm -i hadolint/hadolint < Dockerfile

build:
	@echo Building ${IMAGE_NAME}.
	docker build -t ${IMAGE_NAME} .

test:
	@echo "test: nothing to do for image ${IMAGE_NAME}."

push:
ifeq (${TRAVIS_PULL_REQUEST},false)
	@echo "Pushing image ${IMAGE_NAME} to docker hub."
	@echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
	docker push ${IMAGE_NAME}
else
	@echo "Push: nothing to do for image ${IMAGE_NAME}."
endif

clean:
	@echo "clean: nothing to do for image ${IMAGE_NAME}."

.PHONY: build deps lint test push clean