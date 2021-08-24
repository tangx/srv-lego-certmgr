GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
LAST_COMMIT := $(shell git rev-parse --short HEAD)
VERSION := v$(shell cat .version)

GOMOD := $(shell head -n 1 go.mod | cut -d ' ' -f 2)
FLAGS := "-X $(GOMOD)/version.Version=$(VERSION)-sha.$(LAST_COMMIT)"

MAIN_ROOT := cmd/certmgr

debug:
	cd $(MAIN_ROOT) && go run . --alidns

help:
	cd $(MAIN_ROOT) && go run . --help

build:
	cd $(MAIN_ROOT) && GGO_ENABLED=0 go build -ldflags $(FLAGS) -o ../../bin/certmgr-$(VERSION)-$(GOOS)-$(GOARCH) .

buildx:
	GOOS=darwin GOARCH=amd64 make build
	GOOS=linux  GOARCH=amd64 make build
	GOOS=linux  GOARCH=arm64 make build

clean:
	go mod tidy
	rm -rf bin/
	rm -f cmd/certmgr/*.zip