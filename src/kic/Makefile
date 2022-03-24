.PHONY: all build test

all : cover

build :
	@# build
	@gofmt -s -w .
	@go build -o ../../bin/flt main.go
	@go build -o ../../../fleet-vm/bin/kic main.go
	@go build -o ../../../inner-loop/bin/kic main.go

	# make complete

test : build
	@KIC_PATH=/workspaces/akdc/bin KIC_NAME=flt go test ./... -cover

cover : build
	@KIC_PATH=/workspaces/akdc/bin KIC_NAME=flt go test ./... -cover -coverprofile cover.out
	@go tool cover -html cover.out -o cover.html

	# coverage updated
