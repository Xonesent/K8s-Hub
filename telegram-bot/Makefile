build-server-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o telegram-bot ./cmd/main.go

docker-build:
	docker build -t telegram-bot .

docker-run:
	docker run -d --name telegram-bot --network k8s-hub-network \
 	--env-file .env \
 	-v $(pwd)/config/dev_config.yaml:/config/dev_config.yaml \
 	telegram-bot:latest

proto-gen-statistics:
	protoc --proto_path api \
	--go_out=../statistics-sender/pkg/api/tg_proto --go_opt=paths=source_relative \
	--go-grpc_out=../statistics-sender/pkg/api/tg_proto --go-grpc_opt=paths=source_relative \
	api/tg.proto

proto-gen-telegram-bot:
	protoc --proto_path api \
	--go_out=pkg/api/tg_proto --go_opt=paths=source_relative \
	--go-grpc_out=pkg/api/tg_proto --go-grpc_opt=paths=source_relative \
	api/tg.proto

govulncheck:
	govulncheck ./...