.PHONY: build

build :
	@gofmt -s -w .
	@go build -o ../../bin/kic main.go
	@cp /workspaces/akdc/src/_kic ~/.oh-my-zsh/completions
	# build complete
