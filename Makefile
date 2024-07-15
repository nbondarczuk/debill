TARGET=api

SRC := $(wildcard *.go cmd/*/*.go internal/*/*.go pkg/*/*.go)

VERSION=$(shell git describe --tags --long --dirty 2>/dev/null)

ifeq ($(VERSION),)
    VERSION = UNKNOWN
endif

LDFLAGS=-ldflags "-X main.version=${VERSION}"

# Trigger usage of CGDEBUG env var in docker/image target
GODEBUG=gctrace=1

$(TARGET): $(SRC)
	go build $(LDFLAGS) -o ./bin/$@ ./cmd/$(TARGET)/main.go

schema:
	make -C ./schema init

install:
	go install gorm.io/gen/tools/gentool@latest

gen:
	gentool -c ./schema/gen/gen.tool

clean:
	find . -name *~ -exec rm {} \;
	find . -name .DS_Store -exec rm {} \;
	find . -name *.log -exec rm {} \;

help:
	@echo 'Common build targets'
	@echo
	@echo 'Usage:'
	@echo '    make schema'
	@echo '    make install'
	@echo '    make gen'
	@echo '    make clean'
	@make --no-print-directory go/help test/help pg/help

.PHONY:	$(TARGET) schema install gen clean

include .env
export $(shell sed 's/=.*//' .env)
-include build/include/include.*.mk
