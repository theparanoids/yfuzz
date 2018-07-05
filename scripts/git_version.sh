#!/bin/bash
#
# Copyright 2018 Oath, Inc.
# Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.
#
# Look at the commit message, and build a tag from it.
#

LAST_TAG=$(git describe --tags --abbrev=0)
SEMVER_REGEX='v([0-9]+)\.([0-9]+)\.([0-9]+)'
MAJOR=0
MINOR=0
PATCH=1

if [[ $LAST_TAG =~ $SEMVER_REGEX ]]; then
  MAJOR=${BASH_REMATCH[1]}
	MINOR=${BASH_REMATCH[2]}
	PATCH=${BASH_REMATCH[3]}
  printf "Detected semantic versioning:\nMajor version: %s\nMinor version: %s\nPatch version:%s\n" $MAJOR $MINOR $PATCH
else
	echo "No semantic versioning tag found, defaulting to v0.0.1."
fi

printf "Commit message from Travis:\n%s\n\n" "$TRAVIS_COMMIT_MESSAGE"

case "$TRAVIS_COMMIT_MESSAGE" in
  *[Ss]"emver: "[Mm]"ajor"*)
    echo "Semver: bumping major version to $((++MAJOR))."
		MINOR=0
		PATCH=0
    ;;
  *[Ss]"emver: "[Mm]"inor"*)
    echo "Semver: bumping minor version to $((++MINOR))."
		PATCH=0
    ;;
  *[Ss]"emver: "[Pp]"atch"*)
    echo "Semver: bumping patch version to $((++PATCH))."
    ;;
  *)
    echo 'No version changes found.'
    echo 'Include "Semver: Major|Minor|Patch" in your commit message to bump version.'
esac

export YFUZZ_BUILD_VERSION="v$MAJOR.$MINOR.$PATCH"
echo New version is $YFUZZ_BUILD_VERSION
