#
# SPDX-License-Identifier: BSD-3-Clause
#

PKGS := $(shell go list ./... | grep -v example | grep -v tools)

all: build test

test:
	go test -v $(PKGS)

build:
	go build

clean:
	go clean
