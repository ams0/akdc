.PHONY: build all

build :
	@gofmt -s -w .
	@go build -o ./kic main.go
	# build complete

all : build
	@# copy the cli bin
	@cp kic ${REPO_BASE}/bin
	@cp kic ${REPO_BASE}/src/kic-vm
	@cp kic ${REPO_BASE}/vm-setup-files/akdc/bin

	@# update modules / commands
	@rm -rf ${REPO_BASE}/bin/.kic
	@cp -r ../kic/.kic ${REPO_BASE}/bin

	@# update modules / commands
	@rm -rf ${REPO_BASE}/vm-setup-files/akdc/bin/.kic
	@cp -r ../kic-vm/.kic ${REPO_BASE}/vm-setup-files/akdc/bin

	# copy complete
