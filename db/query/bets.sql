-- name: GetBetsByUser :many
SELECT * FROM bets WHERE user_id = $1;

-- name: GetBetsByMatch :many
SELECT * FROM bets WHERE match_id = $1;

-- name: GetBet :one
SELECT * FROM bets WHERE bet_id = $1;

-- name: CreateBet :one
INSERT INTO bets (user_id, match_id, bet_points, bet_result) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateBet :one
UPDATE bets SET user_id = $2, match_id = $3, bet_points = $4, bet_result = $5 WHERE bet_id = $1 RETURNING *;

-- name: DeleteBet :one
DELETE FROM bets WHERE bet_id = $1 RETURNING *;

-- name: GetBetsByUserAndMatch :many
SELECT * FROM bets WHERE user_id = $1 AND match_id = $2;
