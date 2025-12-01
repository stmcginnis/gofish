#
# SPDX-License-Identifier: BSD-3-Clause
#

PKGS := $(shell go list ./... | grep -v example | grep -v tools)
ROOT_DIR := $(shell git rev-parse --show-toplevel)
GOLANGCI_VERSION := "v2.6.2"

all: lint build test

test:
	go test -v $(PKGS) -cover -race

build:
	go build

lint:
	docker run --rm \
                -v "$(ROOT_DIR)":/src \
                -w /src \
                "golangci/golangci-lint:$(GOLANGCI_VERSION)" \
                golangci-lint run -v

fmt:
	docker run --rm \
                -v "$(ROOT_DIR)":/src \
                -w /src \
                "golangci/golangci-lint:$(GOLANGCI_VERSION)" \
                golangci-lint fmt .


clean:
	go clean
