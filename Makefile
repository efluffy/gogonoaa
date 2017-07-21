# Go makefile adapted from: https://gist.github.com/subfuzion/0bd969d08fe0d8b5cc4b23c795854a13

SHELL := /bin/bash
TARGET := bin/gogonoaa
.DEFAULT_GOAL: $(TARGET)
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
SRC = $(shell find src/ -type f -name '*.go')

.PHONY: all build clean install uninstall check

all: check install

$(TARGET):
	@go build -o $(TARGET) $(LDFLAGS) $(SRC)

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

install:
	@go install $(LDFLAGS) $(SRC)

uninstall: clean
	@rm -f $$(which ${TARGET})

check:
	@test -z $(shell gofmt -l src/main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@go tool vet ${SRC}

