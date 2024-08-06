install:
	@go build -o license-cli
	@mv license-cli $(GOPATH)/bin
	@echo "license-cli installed to $(GOPATH)/bin"