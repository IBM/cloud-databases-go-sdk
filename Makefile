# Makefile to build the project
GOFMT_FILES?= $$(find . -path ./.direnv -prune -false -o -name '*.go')
COVERAGE    = -coverprofile=coverage.txt -covermode=atomic

all: fmt test lint tidy
travis-ci: test-cov lint tidy

fmt:
	gofmt -w $(GOFMT_FILES)

tidy:
	go mod tidy

test:
	go test `go list ./...`

test-cov: 
	go test `go list ./...` ${COVERAGE}

test-int:
	go test `go list ./...` -tags=integration

test-int-cov:
	go test `go list ./...` -tags=integration ${COVERAGE}

lint:
	golangci-lint run
