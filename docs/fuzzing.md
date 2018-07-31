# Using yFuzz
This document will help you get started with [yFuzz](https://github.com/yahoo/yfuzz). It assumes you already have a basic knowledge of fuzzing and building Docker images.

If you've never used a fuzzer before, Google has an excellent tutorial on using [LibFuzzer](https://github.com/google/fuzzer-test-suite/blob/master/tutorial/libFuzzerTutorial.md).

## Table of Contents
- [Using yFuzz](#using-yfuzz)
  - [Table of Contents](#table-of-contents)
  - [Create a Fuzzing Image](#create-a-fuzzing-image)
  - [Add yFuzz Scripts](#add-yfuzz-scripts)
  - [Publish your Image](#publish-your-image)
  - [Create the yFuzz Job](#create-the-yfuzz-job)
  - [List Jobs](#list-jobs)
  - [Get Job Status](#get-job-status)
  - [Retrieve job logs](#retrieve-job-logs)
  - [Delete a Job](#delete-a-job)

## Create a Fuzzing Image
Create a docker image that builds and compiles a fuzz target. For the purposes of this example, this will be an executable located at `/fuzzer`.

## Add yFuzz Scripts
yFuzz contains some shell scripts to organize fuzzing containers and make sure everything is in sync. Add them to your Dockerfile as follows:

```Dockerfile
COPY --from=yfuzz/scripts yfuzz_init.sh /
CMD /yfuzz_init.sh
```

## Publish your Image
In order for Kubernetes to create pods from your image, it needs to be published on an image repository such as [Docker Hub](https://hub.docker.com/).

## Create the yFuzz Job
Install the [yFuzz CLI](../cmd/yfuzz-cli), and ensure it has a `config.yaml` file pointing it to your yFuzz server instance.

To create a job, use the `yfuzz-cli create` command:

```bash
$ yfuzz-cli create https://hub.docker.com/your/image/here
```

## List Jobs
You can see which jobs you have active with the `yfuzz-cli list` command:

```bash
$ yfuzz-cli list
```

## Get Job Status
Let's check on the status of our job:

```bash
$ yfuzz-cli status sample-job
```

## Retrieve job logs
For more detailed information, including data on any crashes that have been found, use `yfuzz-cli logs`:

```bash
$ yfuzz-cli logs sample-job --tail 50
```

## Delete a Job
Once you retrieve the crash, clean up with `yfuzz-cli delete`:

```bash
$ yfuzz-cli delete sample-job
```
