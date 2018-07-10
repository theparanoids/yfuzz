#
# Makefile to recursively test all subprojects of yFuzz with their own Makefiles
#
# Copyright 2018 Oath, Inc.
# Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.
#
# By default it's not very noisy, call with verbose=true for more detailed information.
#

# Information about projects, for calling the proper Makefiles
projects := yfuzz-scripts yfuzz-cli yfuzz-server
yfuzz-scripts-path := images/scripts
yfuzz-cli-path := cmd/yfuzz-cli
yfuzz-server-path := services/yfuzz-server

# Images to be pushed to docker hub
images := scripts

# General information
GIT_SHA := $(shell git rev-parse --short HEAD)

define call_all
	@$(foreach project,${projects},make --directory=${${project}-path} ${1} || STATUS=$$? && echo; && if [[ $${STATUS} -ne 0 ]]; then exit $${STATUS}; fi;)
endef

all: deps lint test

# Wrappers to call some common functions on all projects
deps:
	$(call call_all,deps)

lint:
	$(call call_all,lint)

test:
	$(call call_all,test)

clean:
	$(call call_all,clean)

# Call a makefile of a specific subproject, or a deploy step
travis:
ifeq (${target},deploy-github)
	@echo Tagging ${YFUZZ_BUILD_VERSION} on GitHub
	@git config --global user.email "builds@travis-ci.com"
	@git config --global user.name "Travis CI"
	@git tag -a -m "Generated tag from TravisCI build ${TRAVIS_BUILD_NUMBER}" ${YFUZZ_BUILD_VERSION} 
	@git push https://${GH_TOKEN}@github.com/yahoo/yfuzz.git ${YFUZZ_BUILD_VERSION} > /dev/null 2>&1 
else ifeq (${target},deploy-dockerhub)
	@echo Pushing images to Docker Hub
	@echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
	for image in ${images}; do \
		docker pull yfuzz/$${image}:build-${GIT_SHA}; \
		docker tag yfuzz/$${image}:build-${GIT_SHA} yfuzz/$${image}:latest; \
		docker tag yfuzz/$${image}:build-${GIT_SHA} yfuzz/$${image}:${YFUZZ_BUILD_VERSION}; \
		docker push yfuzz/$${image}; \
	done
else
	make --directory=${${target}-path}
endif

.PHONY: deps lint test clean subproject deploy-github deploy-dockerhub
