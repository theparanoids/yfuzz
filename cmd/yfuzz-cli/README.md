# yFuzz CLI

![godoc](https://godoc.org/github.com/yahoo/yfuzz/cmd/yfuzz-cli?status.svg)

A simple command-line utility for [yFuzz](https://github.com/yahoo/yfuzz).

![Diagram](/yfuzz.png)

## Table of Contents
- [yFuzz CLI](#yfuzz-cli)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites:](#prerequisites)
  - [Install](#install)
  - [Usage](#usage)
    - [Commands](#commands)
  - [Settings](#settings)

## Prerequisites: 
To build the CLI, you will need [Go](https://golang.org/), [Glide](https://glide.sh/), and [Make](https://www.gnu.org/software/make/).

## Install

```
$ git clone https://github.com/yahoo/yfuzz.git
$ cd yfuzz/cmd/yfuzz-cli
$ make install
```

## Usage

```
$ yfuzz-cli [COMMAND]
```

### Commands
* `create, c`: Create a job from a docker image.
* `list, ls`: List all jobs
* `status, st`: Get the status of the specified job
* `delete, d`: Delete the specified job
* `help, h`: Shows a list of commands or help for one command

## Settings
The yFuzz CLI will read configuration from a file called `cli-config.yaml` (or any other format supported by [viper](https://github.com/spf13/viper)) located either in `$HOME/.yfuzz`, `/etc/yfuzz`, or the current directory.

`api`: yFuzz server URL
`athenz.user-cert`: Path to an Athenz user x509 certificate for authentication
`athenz.user-key`: Path to the private key associated with the x509 certificate
`athenz.ca-cert`: CA Cert used by the Athenz server
