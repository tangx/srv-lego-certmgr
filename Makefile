GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
LAST_COMMIT := $(shell git rev-parse --short HEAD)
VERSION := $(shell cat .version)

GOMOD := $(shell head -n 1 go.mod | cut -d ' ' -f 2)
FLAGS := "-X $(GOMOD)/version.Version=$(VERSION)-sha.$(LAST_COMMIT)"

MAIN_ROOT := cmd/certmgr

debug:
	cd $(MAIN_ROOT) && go run .

build:
	cd $(MAIN_ROOT) && go build -ldflags $(FLAGS) -o ../../bin/certmgr-$(GOOS)-$(GOARCH) .

buildx:
	GOOS=darwin GOARCH=amd64 make build
	GOOS=linux  GOARCH=amd64 make build
	GOOS=linux  GOARCH=arm64 make build

clean:
	go mod tidy
	rm -rf bin/