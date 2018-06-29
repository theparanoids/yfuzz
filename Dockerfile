# Copyright 2018 Oath, Inc.
# Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms.

FROM golang:alpine as builder
RUN apk update && apk upgrade && apk add git glide

WORKDIR src/github.com/yahoo/yfuzz
COPY . ./
RUN glide install
RUN go test ./...
RUN go vet ./...
RUN git describe --tags --exact-match `git rev-parse HEAD` | sed 's/^v//' >> version
RUN go install --ldflags "-s -w \
                          -X github.com/yahoo/yfuzz/config.Version=$(cat version)"

FROM alpine as dist
WORKDIR /etc/yfuzz
COPY --from=builder /go/bin/yfuzz ./yfuzz-server
EXPOSE 443
CMD ["/etc/yfuzz/yfuzz-server"]
