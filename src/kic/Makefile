.PHONY: all build test

all : test build

build :
	@# build
	@gofmt -s -w .
	@go build -o ../../bin/kic main.go

	@# copy the cli bin
	@cp ../../bin/kic ../../fleet/bin

	# make complete

test :
	@go test ./...

cover :
	@go test ./... -cover -coverprofile kic.out
	@go tool cover -html kic.out -o kic.html
