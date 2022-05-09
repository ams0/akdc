.PHONY: build

build :
	@# build
	@gofmt -s -w .
	@go build -o ../../bin/flt main.go
	@go build -o ../../vm/bin/kic main.go

	# make complete
