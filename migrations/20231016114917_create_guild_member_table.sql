-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS guild_member (
    id SERIAL PRIMARY KEY,
    guild_id VARCHAR(20) NOT NULL,
    hunter_id VARCHAR(20) NOT NULL,
    nick_name VARCHAR(255) DEFAULT '',
    FOREIGN KEY (guild_id) REFERENCES guild(id) ON DELETE CASCADE,
    FOREIGN KEY (hunter_id) REFERENCES hunter(id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS guild_member;
-- +goose StatementEnd
