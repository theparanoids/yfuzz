# yFuzz Server

![godoc](https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server?status.svg)

The main API server for [yFuzz](https://github.com/yahoo/yfuzz).

## Table of Contents
- [yFuzz Server](#yfuzz-server)
  - [Table of Contents](#table-of-contents)
  - [Configuration](#configuration)
  - [Running](#running)
  - [Plugins](#plugins)
  - [Usage](#usage)
  - [Build](#build)
    - [Local Build](#local-build)
    - [Docker Build](#docker-build)

## Configuration
yFuzz will read configuration from a file called `config.yaml` (or any other format supported by [viper](https://github.com/spf13/viper)) located either in `$HOME/.yfuzz`, `/etc/yfuzz`, or the current directory.

Options can also be specified in environment variables with the `YFUZZ_` prefix.

See `config-sample.yaml` for sample configuration.

## Running
The simplest way to run the server is as a docker container:
```bash
$ docker run -v "$(pwd)"/config.yaml:/etc/yfuzz/config.yaml yfuzz/server
```

## Plugins
A number of plugins to the yFuzz API are supported:
* [Athenz](plugins/athenz): Authorize requests with [Athenz](http://www.athenz.io).
* [MTLS](plugins/mtls): Authenticate requests with mutual TLS and authorize based on a list of authorized keys. 

## Usage
API endpoints are documented with godoc. 

yFuzz is currently accessible through the use of the [yFuzz CLI](../../cmd/yfuzz-cli).

## Build 
To build the server, you will need [Go](https://golang.org/), [Glide](https://glide.sh/), and [Make](https://www.gnu.org/software/make/).

There are two ways to build the yFuzz server: on your system, and as a docker image.

### Local Build
```bash
$ git clone https://github.com/yahoo/yfuzz.git
$ cd yfuzz/services/yfuzz-server
$ make install
```

### Docker Build
```bash
$ git clone https://github.com/yahoo/yfuzz.git
$ cd yfuzz/services/yfuzz-server
$ make docker
```
