export GOPATH := $(PWD):$(PWD)/gopath
export PATH := $(PWD)/bin:$(PWD)/gopath/bin:$(PATH)
export GOMAXPROCS := 1

GOCMD := go

.PHONY: all
all: get capnproto proto capn

.PHONY: proto
proto:
	go version
	go install -v code.google.com/p/goprotobuf/protoc-gen-go
	go install -v code.google.com/p/gogoprotobuf/protoc-gen-gogo
	cd src && protoc --go_out=. pb/log.proto
	cd src && protoc -I$(PWD)/gopath/src -I. --gogo_out=. gogopb/log.proto
	cd src && protoc -I$(PWD)/gopath/src -I. --gogo_out=. gogopb_nullable/log.proto
	cd src && protoc -I$(PWD)/gopath/src -I. --gogo_out=. gogopb_unsafe/log.proto
	cd src && protoc -I$(PWD)/gopath/src -I. --gogo_out=. gogopb_both/log.proto
	go test code.google.com/p/gogoprototest
	go test pb
	go test gogopb gogopb_nullable gogopb_unsafe gogopb_both
	go test -run=XXX -bench=LogProto -benchmem gogopb gogopb_nullable gogopb_unsafe gogopb_both

.PHONY: capn
capn:
	go version
	go install -v github.com/glycerine/go-capnproto/capnpc-go
	capnp compile --verbose -ogo $(PWD)/src/capnp/log.capnp $(PWD)/src/capnp/country.capnp
	go test -c goser
	./goser.test
	./goser.test -test.benchtime=10s -test.cpuprofile=cpu.prof -test.run=XXX -test.bench=. -test.benchmem
	go tool pprof --svg goser.test cpu.prof > cpu.svg

.PHONY: get
get:
	GOPATH=$(PWD)/gopath go get -u -d code.google.com/p/goprotobuf/proto
	GOPATH=$(PWD)/gopath go get -u -d code.google.com/p/gogoprotobuf/proto
	GOPATH=$(PWD)/gopath go get -u -d code.google.com/p/gogoprototest
	GOPATH=$(PWD)/gopath go get -u -d github.com/glycerine/go-capnproto
	GOPATH=$(PWD)/gopath go get -u -d github.com/kaos/capnp_test || true
	GOPATH=$(PWD)/gopath go get code.google.com/p/go.tools/cmd/benchcmp

.PHONY: capnproto
capnproto:
	test -d capnproto || git clone https://github.com/kentonv/capnproto
	cd capnproto/c++ && \
	autoreconf -i && \
	./setup-autotools.sh && \
	./configure --prefix=$(PWD) && \
	make -j4 && \
	make install	

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
