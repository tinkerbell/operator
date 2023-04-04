ARG GO_VERSION=1.20.1
FROM golang:${GO_VERSION} AS builder
WORKDIR /go/src/github.com/moadqasem/kubetink
COPY . .
RUN make all

FROM alpine:3.16

COPY --from=builder \
    /go/src/github.com/moadqasem/kubetink/_build/tinkerbell \
    /usr/local/bin/

USER nobody
