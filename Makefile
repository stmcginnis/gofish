#
# SPDX-License-Identifier: BSD-3-Clause
#

PKGS := $(shell go list ./... | grep -v example | grep -v tools)

all: lint build test

test:
	go test -v $(PKGS)

build:
	go build

lint:
	golangci-lint run -v

clean:
	go clean
