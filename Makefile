.PHONY: proto
proto:
	protoc --proto_path=api/proto --go_out=. api/proto/app/service/*.proto
	protoc --proto_path=api/proto --go_out=. api/proto/app/domain/*.proto
	protoc --proto_path=api/proto --go-grpc_out=. api/proto/app/service/*.proto
	protoc --proto_path=api/proto --grpc-gateway_out=. api/proto/app/service/*.proto

.PHONY: doc
doc:
	protoc --proto_path=api/proto \
           --openapiv2_out=docs \
           --openapiv2_opt use_go_templates=true \
           --openapiv2_opt json_names_for_fields=false \
           --openapiv2_opt allow_delete_body=true \
           --openapiv2_opt visibility_restriction_selectors=PREVIEW \
           --openapiv2_opt proto3_optional_nullable=true \
           --openapiv2_opt disable_default_errors=true \
           api/proto/app/service/*.proto
