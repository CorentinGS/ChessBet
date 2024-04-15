-- name: GetMatches :many
SELECT * FROM matches;

-- name: GetMatch :one
SELECT * FROM matches WHERE match_id = $1;

-- name: CreateMatch :one
INSERT INTO matches (tournament_id, player1_id, player2_id,
 match_date, round_name, lichess_round_id, lichess_game_id, match_result) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: UpdateMatch :one
UPDATE matches SET tournament_id = $2, player1_id = $3, player2_id = $4,
 match_date = $5, round_name = $6, lichess_round_id = $7, lichess_game_id = $8, match_result = $9 WHERE match_id = $1 RETURNING *;
 
-- name: DeleteMatch :one
DELETE FROM matches WHERE match_id = $1 RETURNING *;

-- name: GetMatchesByTournament :many
SELECT * FROM matches WHERE tournament_id = $1;

-- name: GetMatchesByDate :many
SELECT * FROM matches WHERE match_date = $1;

-- name: GetMatchesByPlayer :many
SELECT * FROM matches WHERE player1_id = $1 OR player2_id = $1;

-- name: GetMatchesByPlayerAndDate :many
SELECT * FROM matches WHERE (player1_id = $1 OR player2_id = $1) AND match_date = $2;

-- name: GetMatchesByTournamentAndDate :many
SELECT * FROM matches WHERE tournament_id = $1 AND match_date = $2;

-- name: GetMatchesByTournamentAndPlayer :many
SELECT * FROM matches WHERE tournament_id = $1 AND (player1_id = $2 OR player2_id = $2);

-- name: GetPastMatchesByTournament :many
SELECT * FROM matches WHERE tournament_id = $1 AND match_date < CURRENT_DATE;

-- name: GetUpcomingMatchesByTournament :many
SELECT * FROM matches WHERE tournament_id = $1 AND match_date > CURRENT_DATE;

-- name: GetCurrentMatchesByTournament :many
SELECT * FROM matches WHERE tournament_id = $1 AND match_date = CURRENT_DATE;

-- name: CreateMatches :copyfrom
INSERT INTO matches (tournament_id, player1_id, player2_id, 
match_date, round_name, lichess_round_id, lichess_game_id, match_result) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);