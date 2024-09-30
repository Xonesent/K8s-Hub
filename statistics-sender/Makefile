
build-server-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o sender-statistics-linux ./main.go

docker-build:
	docker build -t sender-statistics .

docker-run:
	docker run -d --name sender-statistics --network k8s-hub-network --env-file .env sender-statistics:latest

govulncheck:
	govulncheck ./...