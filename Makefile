.PHONY: all build test lint

all: build

build:
	@./scripts/build.sh

test:
	@./scripts/test.sh

lint:
	@go vet ./...
