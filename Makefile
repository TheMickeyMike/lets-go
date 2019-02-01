GIT_VERSION := $(shell git describe --abbrev=4 --dirty --always --tags)
GOPATH := $(shell go env GOPATH)

export GO111MODULE=on


build: clean
	go build \
		-ldflags "-X main.appVersion=$(GIT_VERSION)" \
		-o ./build/app cmd/fly/*

clean:
	rm -rf ./build

.PHONY: clean