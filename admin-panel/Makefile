export-env:
	export $(cat .env | xargs)

swagger:
	swag init --parseInternal -g cmd/main.go -g internal/admin_panel/admin_delivery/http/handlers.go -g internal/business_models/admin_panel.go --output cmd/docs

build-server-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o admin-panel-linux ./main.go

docker-build:
	docker build -t admin-panel .

docker-run:
	docker run -d --name admin-panel --network k8s-hub-network --env-file .env -p 8841:8841 admin-panel:latest

proto-gen-statistics:
	protoc --proto_path api \
	--go_out=../statistics-sender/pkg/api/admin_proto --go_opt=paths=source_relative \
	--go-grpc_out=../statistics-sender/pkg/api/admin_proto --go-grpc_opt=paths=source_relative \
	api/admin.proto

proto-gen-admin:
	protoc --proto_path api \
	--go_out=pkg/api/admin_proto --go_opt=paths=source_relative \
	--go-grpc_out=pkg/api/admin_proto --go-grpc_opt=paths=source_relative \
	api/admin.proto

govulncheck:
	govulncheck ./...