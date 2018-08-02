FROM golang:alpine as builder
RUN apk --no-cache add git=2.18.0-r0 glide=0.13.1-r0 make=4.2.1-r2
COPY services /go/src/github.com/yahoo/yfuzz/services
COPY pkg /go/src/github.com/yahoo/yfuzz/pkg
COPY scripts /go/src/github.com/yahoo/yfuzz/scripts
COPY .git /go/src/github.com/yahoo/yfuzz/.git
WORKDIR /go/src/github.com/yahoo/yfuzz/services/yfuzz-server
RUN make deps test install verbose=true

FROM alpine:3.8
COPY --from=builder /go/bin/yfuzz-server /bin/yfuzz-server
CMD ["yfuzz-server"]