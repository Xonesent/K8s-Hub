clickhouse_init:
	docker run -d --name clickhouse-server --network clickhouse-network -p 8123:8123 -p 9000:9000 --ulimit nofile=262144:262144 \
		-e CLICKHOUSE_USER=${CLICKHOUSE_USER} \
        -e CLICKHOUSE_PASSWORD=${CLICKHOUSE_PASSWORD} \
        -e CLICKHOUSE_DB=${CLICKHOUSE_DATABASE} \
        yandex/clickhouse-server

lint:
	golangci-lint -v run ./...

gofumpt:
	gofumpt -l -w .

migrate-up:
	goose -dir ./migrations clickhouse "tcp://${CLICKHOUSE_USER}:${CLICKHOUSE_PASSWORD}@localhost:9000" up
# goose -dir ./migrations clickhouse "tcp://default:fb928275-8771-44e2-a73d-739afa94d725@localhost:9000" up

migrate-down:
	goose -dir ./migrations clickhouse "tcp://${CLICKHOUSE_USER}:${CLICKHOUSE_PASSWORD}@localhost:9000" down
# goose -dir ./migrations clickhouse "tcp://default:fb928275-8771-44e2-a73d-739afa94d725@localhost:9000" down