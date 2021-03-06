GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=taskmgr
LINTER=golangci-lint


all:
	lint test build

test:$(GOTEST) -v

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

lint:
	$(LINTER) run
