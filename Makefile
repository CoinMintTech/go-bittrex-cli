SHELL = /bin/bash
PLATFORM = $(shell go env GOOS)
ARCH = $(shell go env GOARCH)
GOPATH = $(shell go env GOPATH)
GOBIN = $(GOPATH)/bin

default: build

build:
	go fmt ./...
	echo $(PLATFORM)
	echo $(ARCH)
	DEP_BUILD_PLATFORMS=$(PLATFORM) DEP_BUILD_ARCHS=$(ARCH) ./bin/build-all.bash
	cp ./release/bittrex-cli-$(PLATFORM)-$(ARCH) bittrex-cli

install: build
	cp ./bittrex-cli $(GOBIN)

test:
	go test ./...

vendor:
	dep ensure

.PHONY: build test install
