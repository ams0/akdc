.PHONY: build test cover all kubecon

all : cover

build :
	@# build
	@gofmt -s -w .

	@go build -o ../../bin/flt -ldflags="-X 'kic/cmd.TargetCli=flt'" main.go
	@go build -o ../../bin/kic -ldflags="-X 'kic/cmd.TargetCli=kic'" main.go
	@go build -o ../../bin/kivm -ldflags="-X 'kic/cmd.TargetCli=kivm'" main.go

	@go build -o ../../vm/bin/kivm -ldflags="-X 'kic/cmd.TargetCli=kivm'" main.go

	# make complete

test : build
	@go test ./... -cover

cover : build
	@go test ./... -cover -coverprofile cover.out
	@go tool cover -html cover.out -o cover.html

	# coverage updated

kubecon :
	@go build -o ../../kubecon/kic -ldflags="-X 'kic/cmd.TargetCli=kubekic'" main.go
