install:
	@go build -o bin/license-cli
	@cp bin/license-cli $(GOPATH)/bin
	@echo "license-cli installed to $(GOPATH)/bin"