# install:
# 	@go build -o bin/license-cli
# 	@cp bin/license-cli $(GOPATH)/bin
# 	@echo "license-cli installed to $(GOPATH)/bin"

install:
    @CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/license-cli
    @cp bin/license-cli $(GOPATH)/bin
    @echo "license-cli installed to $(GOPATH)/bin"