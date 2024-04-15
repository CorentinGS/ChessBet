-- name: GetPlayers :many
SELECT * FROM players;

-- name: GetPlayer :one
SELECT * FROM players WHERE player_id = $1;

-- name: CreatePlayer :one
INSERT INTO players (name, rating) VALUES ($1, $2) RETURNING *;

-- name: UpdatePlayer :one
UPDATE players SET name = $2, rating = $3 WHERE player_id = $1 RETURNING *;

-- name: DeletePlayer :one
DELETE FROM players WHERE player_id = $1 RETURNING *;
