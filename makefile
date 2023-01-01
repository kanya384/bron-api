.PHONY: help

genproto: ## generates go package from proto files
	@protoc  protos/auth.proto --go-grpc_out=./go/auth --go_out=./go/auth
	@protoc  protos/catalog.proto --go-grpc_out=./go/catalog --go_out=./go/catalog
	@protoc  protos/table.proto --go-grpc_out=./go/table --go_out=./go/table
	@protoc  protos/reserve.proto --go-grpc_out=./go/reserve --go_out=./go/reserve
.PHONY: genproto

gengateway:
	@protoc -I ./protos --openapiv2_out=./openapi --openapiv2_opt allow_merge=true \
   	--go_out ./gateway --go_opt paths=source_relative \
	--go-grpc_out ./gateway --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./gateway --grpc-gateway_opt paths=source_relative \
    ./protos/*.proto

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)