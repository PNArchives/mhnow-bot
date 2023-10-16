-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS painted_monster (
    id SERIAL PRIMARY KEY,
    rank SMALLINT NOT NULL,
    hunter_id VARCHAR(20) NOT NULL,
    monster_id VARCHAR(255) NOT NULL,
    disappear_at TIMESTAMPTZ NOT NULL,
    location VARCHAR(255),
    FOREIGN KEY (hunter_id) REFERENCES hunter(id) ON DELETE CASCADE,
    FOREIGN KEY (monster_id) REFERENCES monster(id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS painted_monster;
-- +goose StatementEnd
