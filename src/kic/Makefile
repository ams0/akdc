.PHONY: build test #all build test

#all : cover

build :
	@# build
	@gofmt -s -w .
	@go build -o /workspaces/akdc/bin/flt -ldflags="-X 'kic/cmd.TargetCli=flt'" main.go
	@go build -o /workspaces/akdc/bin/kic -ldflags="-X 'kic/cmd.TargetCli=kic'" main.go
	@go build -o /workspaces/akdc/bin/kivm -ldflags="-X 'kic/cmd.TargetCli=kivm'" main.go
	@go build -o /workspaces/akdc/vm/bin/kivm -ldflags="-X 'kic/cmd.TargetCli=kivm'" main.go
	@go build -o /workspaces/akdc/kubecon/kic -ldflags="-X 'kic/cmd.TargetCli=kubekic'" main.go

	# make complete

test : build
	@go test ./... -cover

cover : build
	@go test ./... -cover -coverprofile cover.out
	@go tool cover -html cover.out -o cover.html

	# coverage updated
