.PHONY: build test lint run clean install

VERSION ?= dev
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE    ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -ldflags "-X github.com/apzelos/zw/internal/cli.Version=$(VERSION) -X github.com/apzelos/zw/internal/cli.Commit=$(COMMIT) -X github.com/apzelos/zw/internal/cli.Date=$(DATE)"

build:
	go build $(LDFLAGS) -o bin/zw ./cmd/zw

test:
	go test -v ./...

lint:
	golangci-lint run

run:
	go run ./cmd/zw

clean:
	rm -rf bin/

install: build
	cp bin/zw $(GOPATH)/bin/zw
