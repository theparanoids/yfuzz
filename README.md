# YFuzz

YFuzz is a project for running fuzzing jobs at scale with Kubernetes.

## Table of Contents

- [Background](#background)
- [Install](#install)
- [Configuration](#configuration)
- [Usage](#usage)
- [Architecture](#architecture)
- [Contribute](#contribute)
- [License](#license)

## Background

Popular fuzzers such as [Libfuzzer](https://llvm.org/docs/LibFuzzer.html) and [AFL](http://lcamtuf.coredump.cx/afl/) have support for running multiple fuzzing processes at once. YFuzz aims to take advantage of this by running each process on a different Kubernetes pod to speed up the fuzzing process.

## Install

```
$ git clone https://github.com/yahoo/yfuzz.git
$ cd yfuzz
$ glide update && GOBIN=$PWD go install -v github.com/yahoo/yfuzz
```

## Configuration
YFuzz will read configuration from a file called `config.yaml` (or any other format supported by [viper](https://github.com/spf13/viper)) located either in `$HOME/.yfuzz`, `/etc/yfuzz`, or the current directory.

Options can also be specified in environment variables with the `YFUZZ_` prefix.

See `config-sample.yaml` for sample configuration.

## Usage
API endpoints are documented with godoc. 

YFuzz is currently accessible through the use of the [YFuzz CLI](https://github.com/yahoo/yfuzz-cli). A web interface is also planned.

## Architecture
![Architecture Diagram](architecture.png)

The YFuzz API resides in a kubernetes cluster along with the pods that run the fuzzing jobs and a shared volume that holds corpus data to be shared between the pods.

### Planned:
* Each fuzzing pod will have a logging sidecar which streams logs from the pod to a centralized logging service.
* The YFuzz API will have access to a data store with information about users, jobs, and crash files.

## Contribute

Please refer to [the contributing.md file](CONTRIBUTING.md) for information about how to get involved. We welcome issues, questions, and pull requests. Pull Requests are welcome

## License
This project is licensed under the terms of the [Apache 2.0](LICENSE) open source license.
