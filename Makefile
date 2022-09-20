.PHONY: proto
proto:
	protoc --proto_path=api/proto --go_out=. api/proto/service/*.proto
	protoc --proto_path=api/proto --go_out=. api/proto/domain/*.proto
	protoc --proto_path=api/proto --go-grpc_out=. api/proto/service/*.proto
	protoc --proto_path=api/proto --grpc-gateway_out=. api/proto/service/*.proto