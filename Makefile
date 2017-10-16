# Makefile for Go-Project


# dynamically fetch path to executables
GO_BIN := $(GOPATH)/bin
GOMETALINTER := $(GO_BIN)/gometalinter

# in case gometalinter is not installed already => clone it and install it
$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

# fire up gometalinter to concurrently run several static analysis tools at once
# it's PHONY, as it doesn't create a (target-)file
.PHONY: lint
lint: $(GOMETALINTER)
	# recursevly run gometalinter on all files in this directory, skipping packages in vendor
	gometalinter ./... --vendor

# build everything in this directory and sub-directories
# TODO: work this out
.PHONY: build
build:
	go build -o bin/main

# build docker image
# TODO: enhance
.PHONY: dockerbuild
dockerbuild:
	sudo docker build -t test .

.PHONY: dockerrun
dockerrun:
	sudo docker run -P --name test --rm test

# run main service package
.PHONY: run
run:
	go run ./main/*.go

