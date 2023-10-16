-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS monster (
    id VARCHAR(255) PRIMARY KEY,
    jp_name VARCHAR(255) NOT NULL,
    en_name VARCHAR(255) NOT NULL,
    discoverable BOOLEAN NOT NULL DEFAULT TRUE
);
INSERT INTO monster(id, discoverable, en_name, jp_name) VALUES
    ('greatest_jagras', TRUE , 'Greatest Jagras',   'ドスジャグラス'),
    ('kuluyaku',        TRUE , 'Kuluyaku',          'クルルヤック'),
    ('pukeipukei',      TRUE , 'Pukei Pukei',       'プケプケ'),
    ('barroth',         TRUE , 'Barroth',           'ボルボロス'),
    ('great_girros',    TRUE , 'Great Girros',      'ドスギルオス'),
    ('tobikadachi',     TRUE , 'Tobi-Kadachi',      'トビカガチ'),
    ('paolumu',         TRUE , 'Paolumu',           'パオウルムー'),
    ('jyuratodus',      TRUE , 'Jyuratodus',        'ジュラトドス'),
    ('anjanath',        TRUE , 'Anjanath',          'アンジャナフ'),
    ('rathian',         TRUE , 'Rathian',           'リオレイア'),
    ('legiana',         TRUE , 'Legiana',           'レイギエナ'),
    ('diablos',         TRUE , 'Diablos',           'ディアブロス'),
    ('rathalos',        TRUE , 'Rathalos',          'リオレウス'),
    ('black_diablos',   FALSE, 'Black Diablos',     'ディアブロス亜種'),
    ('pink_rathian',    FALSE, 'Pink Rathian',      'リオレイア亜種');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS monster;
-- +goose StatementEnd
