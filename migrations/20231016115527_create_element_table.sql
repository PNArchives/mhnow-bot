-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS element (
    id VARCHAR(20) PRIMARY KEY,
    jp_name VARCHAR(20) NOT NULL,
    en_name VARCHAR(20) NOT NULL
);
INSERT INTO element(id, en_name, jp_name) VALUES
    ('dragon',  'dragon',  '龍'),
    ('fire',    'fire',    '火'),
    ('ice',     'ice',     '氷'),
    ('mahi',    'poison',  '麻痺'),
    ('poison',  'poison',  '毒'),
    ('thunder', 'thunder', '雷'),
    ('water',   'water',   '水');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS element;
-- +goose StatementEnd
