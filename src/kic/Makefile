.PHONY: all test

all :
	@# build
	@gofmt -s -w .
	@go build -o ./kic main.go

	@# copy the cli bin
	@cp kic ../../bin
	@cp kic ../../fleet/bin

	@# update commands
	@rm -rf ../../bin/.kic
	@cp -r .kic ../../bin

	# make complete

test :
	@go test ./... -cover
