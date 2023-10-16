-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS hunter (
    id VARCHAR(20) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    display_name VARCHAR(255) DEFAULT ''
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS hunter;
-- +goose StatementEnd
