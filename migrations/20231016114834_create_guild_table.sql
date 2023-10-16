-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS guild (
    id VARCHAR(20) PRIMARY KEY,
    name VARCHAR(255) DEFAULT ''
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS guild;
-- +goose StatementEnd
