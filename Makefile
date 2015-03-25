export GOPATH := $(PWD):$(PWD)/gopath
export PATH := $(PWD)/gopath/bin:$(PATH)
export GOMAXPROCS := 1

GOCMD := go
CAPNP_VERSION := 0.5.1.2
CAPNP_NAME := capnproto-c++-$(CAPNP_VERSION)
CAPNP_CMD := bin/capnp
PROTOC_VERSION := 2.6.1
PROTOC_NAME := protobuf-$(PROTOC_VERSION)
PROTOC_CMD := bin/protoc

.PHONY: all
all: get proto capn
	go test -c goser
	./goser.test
	./goser.test -test.benchtime=10s -test.cpuprofile=cpu.prof -test.run=XXX -test.bench=. -test.benchmem
	go tool pprof --svg goser.test cpu.prof > cpu.svg

.PHONY: proto
proto: $(PROTOC_CMD)
	go version
	go install -v github.com/golang/protobuf/protoc-gen-go
	go install -v github.com/gogo/protobuf/protoc-gen-gogo
	cd src && ../$(PROTOC_CMD) --go_out=. pb/log.proto
	cd src && ../$(PROTOC_CMD) -I$(PWD)/gopath/src -I$(PWD)/gopath/src/github.com/gogo/protobuf/protobuf -I. --gogo_out=. gogopb/log.proto
	cd src && ../$(PROTOC_CMD) -I$(PWD)/gopath/src -I$(PWD)/gopath/src/github.com/gogo/protobuf/protobuf -I. --gogo_out=. gogopb_nullable/log.proto
	cd src && ../$(PROTOC_CMD) -I$(PWD)/gopath/src -I$(PWD)/gopath/src/github.com/gogo/protobuf/protobuf -I. --gogo_out=. gogopb_unsafe/log.proto
	cd src && ../$(PROTOC_CMD) -I$(PWD)/gopath/src -I$(PWD)/gopath/src/github.com/gogo/protobuf/protobuf -I. --gogo_out=. gogopb_both/log.proto

$(PROTOC_CMD):
	test -d $(PROTOC_NAME) || curl -s -L https://github.com/google/protobuf/releases/download/v$(PROTOC_VERSION)/$(PROTOC_NAME).tar.bz2 | tar jx
	cd $(PROTOC_NAME) && \
	./configure --prefix=$(PWD) && \
	make -j3 && \
	make install

.PHONY: capn
capn: $(CAPNP_CMD)
	go version
	go install -v github.com/glycerine/go-capnproto/capnpc-go
	$(CAPNP_CMD) compile --verbose -ogo $(PWD)/src/capnp/log.capnp $(PWD)/src/capnp/country.capnp

$(CAPNP_CMD):
	test -d $(CAPNP_NAME) || curl -s -L https://capnproto.org/$(CAPNP_NAME).tar.gz | tar zx
	cd $(CAPNP_NAME) && \
	./configure --prefix=$(PWD) && \
	make -j3 && \
	make install

.PHONY: get
get:
	GOPATH=$(PWD)/gopath go get -u -d github.com/golang/protobuf/proto
	GOPATH=$(PWD)/gopath go get -u -d github.com/gogo/protobuf/proto
	GOPATH=$(PWD)/gopath go get -u -d github.com/gogo/harmonytests
	GOPATH=$(PWD)/gopath go get -u -d github.com/glycerine/go-capnproto
	GOPATH=$(PWD)/gopath go get -u -d github.com/kaos/capnp_test || true
	GOPATH=$(PWD)/gopath go get golang.org/x/tools/cmd/benchcmp

.PHONY: gofmt
gofmt:
	$(GOCMD) fmt ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: test-compile
test-compile: all
	@$(MAKE) --no-print-directory -f Make.tests $@

.PHONY: clean-pkg
clean-pkg:
	$(RM) -rf pkg gopath/pkg

.PHONY: clean
clean: clean-pkg
	$(RM) -rf bin gopath/bin

DEFAULT_GOAL := all
