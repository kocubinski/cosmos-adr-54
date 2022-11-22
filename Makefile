GOLANG_PROTOBUF_VERSION = 1.28.1
GRPC_GATEWAY_VERSION = 1.16.0
BIN = $(CURDIR)/bin
BUF = $(BIN)/buf

.PHONY: proto-gen

$(BIN)/protoc-gen-go:
	GOBIN=$(BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v$(GOLANG_PROTOBUF_VERSION)

$(BIN)/protoc-gocosmos:
	GOBIN=$(BIN) go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest

$(BIN)/protoc-gen-grpc-gateway:
	GOBIN=$(BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v$(GRPC_GATEWAY_VERSION)

$(BUF):
	./install-buf.sh

proto-gen: $(BUF) $(BIN)/protoc-gen-go $(BIN)/protoc-gocosmos $(BIN)/protoc-gen-grpc-gateway
	@echo "Generating proto files"
	@buf generate

clean:
	rm -rf $(BIN)