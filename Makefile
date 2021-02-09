GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

MAIN_ROOT := cmd/certmgr

debug:
	cd $(MAIN_ROOT) && go run .

build:
	cd $(MAIN_ROOT) && go build -o ../../bin/certmgr-$(GOOS)-$(GOARCH) .

buildx:
	GOOS=darwin GOARCH=amd64 make build
	GOOS=linux  GOARCH=amd64 make build
	GOOS=linux  GOARCH=arm64 make build

clean:
	go mod tidy
	rm -rf bin/