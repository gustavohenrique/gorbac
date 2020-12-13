.PHONY: test

GO := $(shell which go)

test: tests
tests:
	$(GO) test -v -failfast -cover ./...

lint: linter
linter:
	goimports -w .

install: setup
setup:
	$(GO) get -u golang.org/x/tools/cmd/goimports
