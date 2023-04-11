ARG GO_VERSION=1.20.1
FROM golang:${GO_VERSION} AS builder
WORKDIR /go/src/github.com/tinkerbell/operator
COPY . .
RUN make all

FROM alpine:3.17

COPY --from=builder \
    /go/src/github.com/tinkerbell/operator/_build/tinkerbell \
    /usr/local/bin/

USER nobody
