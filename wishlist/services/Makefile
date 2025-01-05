PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GO_GRPC = protoc-gen-go-grpc
PROTO_FILES = wishlist-service.proto
PROTO_DIR = proto
OUT_DIR_SERVER = wishlist-service
OUT_DIR_CLIENT = frontend-service

.PHONY: generate

generate:
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR_SERVER) --go-grpc_out=$(OUT_DIR_SERVER) \
		$(PROTO_FILES)
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR_CLIENT) --go-grpc_out=$(OUT_DIR_CLIENT) \
		$(PROTO_FILES)


