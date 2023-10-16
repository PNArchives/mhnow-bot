-- name: ListDiscoverableMonsters :many
SELECT *
FROM monster
WHERE discoverable = TRUE;

-- name: RegisterHunter :one
INSERT INTO hunter (
    id,
    name,
    display_name
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetHunter :one
SELECT *
FROM hunter
WHERE id = $1;

-- name: RegisterGuild :one
INSERT INTO guild (
    id,
    name
) VALUES (
    $1, $2
)
RETURNING *;

-- name: JoinGuild :exec
INSERT INTO guild_member (
    guild_id,
    hunter_id,
    nick_name
) VALUES (
    $1, $2, $3
);

-- name: FindHuntableMontersWithGuild :many
SELECT
    pm.rank,
    pm.disappear_at,
    pm.location,
    h.id as hunter_id,
    h.name as hunter_name,
    m.id as monster_id,
    m.jp_name as monster_jp_name,
    m.en_name as monster_en_name
FROM painted_monster pm
INNER JOIN hunter h ON pm.hunter_id = h.id
INNER JOIN monster m ON pm.monster_id = m.id
WHERE
    pm.disappear_at > CURRENT_TIMESTAMP
ORDER BY pm.disappear_at;

-- name: PaintMonster :one
INSERT INTO painted_monster (
    rank,
    hunter_id,
    monster_id,
    disappear_at,
    location
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;
