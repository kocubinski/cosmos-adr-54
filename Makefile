# proto dependencies sourced roughly from
# https://github.com/cosmos/cosmos-sdk/blob/76be73022ac811b97ad4e3f8cf14a41563e795df/contrib/devtools/Dockerfile

GOLANG_PROTOBUF_VERSION = 1.28.1
GRPC_GATEWAY_VERSION = 1.16.0
BIN = $(CURDIR)/bin
BUF = $(BIN)/buf

.PHONY: proto-gen

$(BIN)/protoc-gen-go:
	GOBIN=$(BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v$(GOLANG_PROTOBUF_VERSION)

$(BIN)/protoc-gen-gocosmos:
	GOBIN=$(BIN) go install github.com/cosmos/gogoproto/protoc-gen-gocosmos@latest

$(BIN)/protoc-gen-grpc-gateway:
	GOBIN=$(BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v$(GRPC_GATEWAY_VERSION)

$(BUF):
	./install-buf.sh

proto-gen: $(BUF) $(BIN)/protoc-gen-go $(BIN)/protoc-gen-gocosmos $(BIN)/protoc-gen-grpc-gateway
	@echo "Generating proto files"
	PATH=$(BIN) buf generate --output src/naive/foo/v1/api src/naive/foo/v1/proto

clean:
	rm -rf $(BIN)