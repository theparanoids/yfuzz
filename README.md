# YFuzz

[![Build Status](https://travis-ci.org/yahoo/yfuzz.svg?branch=master)](https://travis-ci.org/yahoo/yfuzz)

YFuzz is a project for running fuzzing jobs at scale with Kubernetes.

## Table of Contents

- [YFuzz](#yfuzz)
  - [Table of Contents](#table-of-contents)
  - [Background](#background)
  - [Projects](#projects)
  - [Architecture](#architecture)
    - [Planned:](#planned)
  - [Directory Structure](#directory-structure)
  - [Contribute](#contribute)
  - [License](#license)

## Background

Popular fuzzers such as [Libfuzzer](https://llvm.org/docs/LibFuzzer.html) and [AFL](http://lcamtuf.coredump.cx/afl/) have support for running multiple fuzzing processes at once. YFuzz aims to take advantage of this by running each process on a different Kubernetes pod to speed up the fuzzing process.

## Projects
* [YFuzz Server](services/yfuzz-server): The main API server for YFuzz.
* [YFuzz CLI](cmd/yfuzz-cli): A command-line interface for interacting with the YFuzz server.

## Architecture
![Architecture Diagram](architecture.png)

The YFuzz API resides in a kubernetes cluster along with the pods that run the fuzzing jobs and a shared volume that holds corpus data to be shared between the pods.

### Planned:
* Each fuzzing pod will have a logging sidecar which streams logs from the pod to a centralized logging service.
* The YFuzz API will have access to a data store with information about users, jobs, and crash files.

## Directory Structure
* `cmd`: Command line utilities.
* `pkg`: Shared libraries and packages.
* `scripts`: Scripts for CI tooling.
* `services`: Long-running services, such as the yfuzz-server.

## Contribute

Please refer to [the contributing.md file](CONTRIBUTING.md) for information about how to get involved. We welcome issues, questions, and pull requests. Pull Requests are welcome

## License
This project is licensed under the terms of the [Apache 2.0](LICENSE) open source license.
