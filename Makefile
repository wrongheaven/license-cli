build:
	@CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/license-cli

install: build
	@cp bin/license-cli $(GOPATH)/bin
	@echo "license-cli installed to $(GOPATH)/bin"