GO ?= go

BINARY := bin/uud
PKG := ./cmd/uud

.PHONY: all build clean

all: build

build: $(BINARY)

$(BINARY): $(shell find cmd pkg -type f -name '*.go') go.mod
	mkdir -p $(dir $@)
	$(GO) build -o $@ $(PKG)

clean:
	rm -f $(BINARY)

