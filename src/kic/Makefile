.PHONY: all build test

all : cover

build :
	@# build
	@gofmt -s -w .
	@go build -o ../../bin/akdc main.go

	# make complete

test : build
	@go test ./... -cover

cover : build
	@go test ./... -cover -coverprofile cover.out
	@go tool cover -html cover.out -o cover.html

	# coverage updated
