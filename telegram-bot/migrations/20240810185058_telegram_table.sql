-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS telegram.users (
    tg_id UInt32,
    message String
) ENGINE = MergeTree()
ORDER BY tg_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP DATABASE IF EXISTS telegram.users;
-- +goose StatementEnd