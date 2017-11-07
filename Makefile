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
	gometalinter ./... --vendor --disable=gotype

# build everything in this directory into a single binary in bin-directory
.PHONY: build
build:
ifeq ($(OS),Windows_NT)
	go build -o bin/main.exe
else
	go build -o bin/main
endif

# build docker image
.PHONY: dockerbuild
dockerbuild:
ifeq ($(OS),Windows_NT)
	docker build -t test .
else
	sudo docker build -t test .
endif


.PHONY: dockerrun
dockerrun:
ifeq ($(OS),Windows_NT)
	docker run -P --name test --rm test
else
	sudo docker run -P --name test --rm test
endif