# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_NAME = git-com
BINARY_UNIX = $(BINARY_NAME)_unix

all: test build
.PHONY: all

build:
	$(GOBUILD) -o $(BINARY_NAME) -v
.PHONY: build

test:
	$(GOTEST) -v ./...
.PHONY: test

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
.PHONY: clean

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
.PHONY: run

deps:
	$(GOGET) -v ./...
.PHONY: deps

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
.PHONY: build-linux
