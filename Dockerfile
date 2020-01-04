FROM golang:latest AS build-env

ENV GO111MODULE on

WORKDIR /go/src/github.com/snaka/go-whatsmyip
