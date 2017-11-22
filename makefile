all: build
build:
	GOPATH=$(shell pwd) gofmt -w src/main src/lib/*
	GOPATH=$(shell pwd) go build -o pcapagent main

clean:
	rm -f pcapagent
