NAME := whatsmyip
SRCS := $(shell find . -type f -name '*.go')

export GO111MODULE=on

## Install dependencies
.PHONY: deps
deps:
	go get -v -d

## Install dependencies (used in development task only)
.PHONY: devel-deps
devel-deps: deps
	GO111MODULE=off go get \
		golang.org/x/lint/golint

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

## build binary if source code updaded
bin/%: main.go lib/%/*.go deps
	go build -v -o $@ $<

## build
.PHONY: build
build: bin/$(NAME)
