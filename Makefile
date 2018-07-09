#
# Makefile to recursively test all subprojects of YFuzz with their own Makefiles
#
# Copyright 2018 Oath, Inc.
# Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.
#
# By default it's not very noisy, call with verbose=true for more detailed information.
#

# Information about projects, for calling the proper Makefiles
projects := yfuzz-scripts yfuzz-cli yfuzz-server
yfuzz-scripts-path := images/yfuzz-scripts
yfuzz-cli-path := cmd/yfuzz-cli
yfuzz-server-path := services/yfuzz-server


define call_all
	@$(foreach project,${projects},make --directory=${${project}-path} ${1}; echo;)
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

# Call a makefile of a specific subproject
subproject:
	make --directory=${${target}-path}

.PHONY: deps lint test clean subproject
