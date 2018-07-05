# YFuzz Server

The main API server for [YFuzz](https://github.com/yahoo/yfuzz).

## Table of Contents
- [YFuzz Server](#yfuzz-server)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites:](#prerequisites)
  - [Install](#install)
  - [Configuration](#configuration)
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
YFuzz will read configuration from a file called `config.yaml` (or any other format supported by [viper](https://github.com/spf13/viper)) located either in `$HOME/.yfuzz`, `/etc/yfuzz`, or the current directory.

Options can also be specified in environment variables with the `YFUZZ_` prefix.

See `config-sample.yaml` for sample configuration.

## Usage
API endpoints are documented with godoc. 

YFuzz is currently accessible through the use of the [YFuzz CLI](../../cmd/yfuzz-cli). A web interface is also planned.