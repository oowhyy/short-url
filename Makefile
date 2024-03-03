testcoverage:
	@go test -coverprofile=c.out ./...
	@go tool cover -func=c.out
	@go tool cover -html=c.out

.PHONY: proto
proto:
	@protoc  --go_out=. --go-grpc_out=. api/v1/shorturl/urls.proto