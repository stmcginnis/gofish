#
# SPDX-License-Identifier: BSD-3-Clause
#

PKGS := $(shell go list ./school/...)

all: build test

test:
	go test -v $(PKGS)

build:
	go build

clean:
	go clean
