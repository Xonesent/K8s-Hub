#!/bin/sh

until curl -s "http://$CLICKHOUSE_HOST:$CLICKHOUSE_PORT/ping" | grep -q "Ok"; do
  echo "Waiting for ClickHouse at $CLICKHOUSE_HOST:$CLICKHOUSE_PORT..."
  sleep 2
done

echo "ClickHouse is available! Proceeding with migrations..."