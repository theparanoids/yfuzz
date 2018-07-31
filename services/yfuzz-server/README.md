# yFuzz Server

![godoc](https://godoc.org/github.com/yahoo/yfuzz/services/yfuzz-server?status.svg)

The main API server for [yFuzz](https://github.com/yahoo/yfuzz).

## Table of Contents
- [yFuzz Server](#yfuzz-server)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites:](#prerequisites)
  - [Install](#install)
  - [Configuration](#configuration)
  - [Plugins](#plugins)
  - [Usage](#usage)

## Prerequisites: 
To build the server, you will need [Go](https://golang.org/), [Glide](https://glide.sh/), and [Make](https://www.gnu.org/software/make/).

## Install

```
$ git clone https://github.com/yahoo/yfuzz.git
$ cd yfuzz/services/yfuzz-server
$ make install
```

## Configuration
yFuzz will read configuration from a file called `config.yaml` (or any other format supported by [viper](https://github.com/spf13/viper)) located either in `$HOME/.yfuzz`, `/etc/yfuzz`, or the current directory.

Options can also be specified in environment variables with the `YFUZZ_` prefix.

See `config-sample.yaml` for sample configuration.

## Plugins
A number of plugins to the yFuzz API are supported:
* [Athenz](plugins/athenz): Authorize requests with [Athenz](http://www.athenz.io).
* [MTLS](plugins/mtls): Authenticate requests with mutual TLS and authorize based on a list of authorized keys. 

## Usage
API endpoints are documented with godoc. 

yFuzz is currently accessible through the use of the [yFuzz CLI](../../cmd/yfuzz-cli).
