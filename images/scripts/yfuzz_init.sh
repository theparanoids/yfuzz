#!/bin/bash
#
# Copyright 2018 Oath, Inc.
# Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.
#

# Check that the variables are set
if ! [[ -n "$YFUZZ_PROJECT" && -n "$YFUZZ_TARGET" ]]; then
  echo \$YFUZZ_PROJECT and \$YFUZZ_TARGET are required.
  exit 1
fi

TARGET_DIR=/shared_data/$YFUZZ_PROJECT/$YFUZZ_TARGET
export CORPUS_DIR=$TARGET_DIR/corpus
export CRASH_FILE=$TARGET_DIR/crash

# Make the corpus dir and add any seed inputs
mkdir -p $CORPUS_DIR
if [[ -n "$SEED_CORPUS_DIR" ]]; then
  cp -r $SEED_CORPUS_DIR/* $CORPUS_DIR
fi

# run fuzzer, then monitor for crashes
/run_fuzzer.sh &
FUZZ_PROCESS=$!

# watch for a local crash, a remote crash, or the process to exit
while [[ ! -f $CRASH_FILE ]] && kill -s 0 $FUZZ_PROCESS 2>/dev/null; do
  sleep 5
done

if kill -s 0 $FUZZ_PROCESS 2>/dev/null; then
  # Process was still running. Somebody else found a crash.
  kill $FUZZ_PROCESS
  echo "Other pod found crash. Exiting."
  exit 0
else
  # Process exited by itself
  wait $FUZZ_PROCESS
  FUZZ_PROCESS_EXITCODE=$?
  echo "Fuzz process exited with code $FUZZ_PROCESS_EXITCODE."
  exit $FUZZ_PROCESS_EXITCODE
fi
