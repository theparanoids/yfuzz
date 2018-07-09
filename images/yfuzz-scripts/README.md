# YFuzz Scripts

Image containing scripts used by containers in [YFuzz](https://github.com/yahoo/yfuzz).

## Table of Contents
- [YFuzz Scripts](#yfuzz-scripts)
  - [Table of Contents](#table-of-contents)
  - [Usage](#usage)
    - [Environment Variables](#environment-variables)

## Usage

Create an image that runs a fuzz target (e.g. with [Libfuzzer](https://llvm.org/docs/LibFuzzer.html)), and add the `yfuzz_init` script:

```Dockerfile
COPY --from=yfuzz/scripts yfuzz_init.sh /
CMD /yfuzz_init.sh
```

### Environment Variables

Variables needed by the script:
* `YFUZZ_PROJECT`: Name of the YFuzz Project  
* `YFUZZ_TARGET`: Name of the target within the YFuzz project
* `FUZZER_COMMAND`: Command to run the fuzzer.

Optional variables:
* `SEED_CORPUS_DIR`: Path to a folder of seed inputs to be copied to the corpus.

YFuzz will set the following environment variables for use by the fuzzer:
* `CORPUS_DIR`: Path to a directory for corpus files to be shared between pods.
* `CRASH_FILE`: Location to write a file with information about a crash, when one is found.
