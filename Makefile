say:
	@echo hello

.PHONY: proto
proto:
	@protoc  --go_out=. --go-grpc_out=. api/v1/shorturl/urls.proto