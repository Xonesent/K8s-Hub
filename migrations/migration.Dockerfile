FROM golang:1.22 AS builder

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /app

RUN mkdir migrations

COPY ./migrations /app/migrations
COPY ./migrations/wait-for-clickhouse.sh /app/wait-for-clickhouse.sh

FROM ubuntu:latest

WORKDIR /app

RUN mkdir migrations

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/wait-for-clickhouse.sh /app/wait-for-clickhouse.sh

RUN apt-get update && apt-get install -y curl && rm -rf /var/lib/apt/lists/*

RUN chmod +x /app/wait-for-clickhouse.sh

CMD ["/bin/sh", "-c", "/app/wait-for-clickhouse.sh && goose -dir /app/migrations clickhouse $CLICKHOUSE_DSN up"]