-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS telegram.messages (
    tg_id UInt32,
    chat_id UInt32,
    message String CODEC(ZSTD),
    created_at DateTime CODEC(T64)
) ENGINE = MergeTree()
ORDER BY tg_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS telegram.messages;
-- +goose StatementEnd