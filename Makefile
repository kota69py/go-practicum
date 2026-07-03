.DEFAULT_GOAL := all

VERSION := $(shell git describe --tags --always --dirty 2>nul || echo dev)
LDFLAGS := -ldflags="-X github.com/kota69py/go-practicum/cmd.version=$(VERSION)"

.PHONY: all build test vet lint clean doctor coverage vulncheck release-dry-run

all: vet test build

build:
	go build $(LDFLAGS) -o go-practicum.exe .

test:
	go test ./... -count=1 -timeout 10m -v

vet:
	go vet ./...

lint:
	golangci-lint run ./...

clean:
	go clean -testcache

doctor:
	go run . doctor

coverage:
	go test ./... -count=1 -timeout 10m -coverprofile=coverage.out -covermode=atomic
	go tool cover -html=coverage.out -o coverage.html

vulncheck:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

release-dry-run:
	goreleaser release --snapshot --clean
