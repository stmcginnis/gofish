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

OUTPUT_DIR ?= ../..
# Schema bundles the committed code is generated from. Bump these to update.
REDFISH_TAG ?= 2026.2
SWORDFISH_TAG ?= v1.2.9_Release

generate:
	cd tools/generator && \
    go run ./cmd/generate-schemas --output-dir $(OUTPUT_DIR) \
		--redfish-tag $(REDFISH_TAG) --swordfish-tag $(SWORDFISH_TAG)

# Fails if regenerating changes or adds any file under schemas/ or serviceroot.go.
check-generate: generate
	git diff --exit-code -- schemas serviceroot.go
	test -z "$$(git status --porcelain -- schemas serviceroot.go)"

clean:
	go clean
