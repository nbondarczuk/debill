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

install:
	go install gorm.io/gen/tools/gentool@latest

db: db/create db/start db/status

schema: schema/init

gen: schema/gen

clean:
	find . -name *~ -exec rm {} \;
	find . -name .DS_Store -exec rm {} \;
	find . -name *.log -exec rm {} \;

purge: db/stop db/drop db/clean

help:
	@echo 'Common build targets'
	@echo
	@echo 'Usage:'
	@echo '    make install - install required binaries'
	@echo '    make db      - install Postgres DB locally'
	@echo '    make schema  - create shema in the local Postgres DB'
	@echo '    make gen     - generate model using local Postgres DB tables'
	@echo '    make clean   - clean files'
	@echo '    make purge   - stop and destroy local Postgres DB'
	@make --no-print-directory go/help test/help db/help schema/help

.PHONY:	$(TARGET) schema install gen clean schema schema/init schema/gen

include .env
export $(shell sed 's/=.*//' .env)
-include build/include/include.*.mk
