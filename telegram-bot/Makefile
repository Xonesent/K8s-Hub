run-without-compile:
	go run cmd/main.go

build-server-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o server-windows.exe ./main.go

build-server-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o server-linux ./main.go

run-after-compile:
	./server-windows.exe

docker-build:
	docker build -t telegram-bot .

docker-run:
	docker run -d --name telegram-bot --network clickhouse-network telegram-bot:latest

govulncheck:
	govulncheck ./...