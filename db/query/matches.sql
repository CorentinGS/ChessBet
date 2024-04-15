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

-- name: GetUpcomingMatchesByTournament :many
SELECT matches.*, player1.*, player2.*
FROM matches 
JOIN players AS player1 ON matches.player1_id = player1.player_id
JOIN players AS player2 ON matches.player2_id = player2.player_id
WHERE matches.tournament_id = $1 AND matches.match_date >= NOW()
GROUP BY matches.match_id, matches.lichess_round_id, player1.player_id, player2.player_id
ORDER BY matches.match_date ASC;

-- name: CreateMatches :copyfrom
INSERT INTO matches (tournament_id, player1_id, player2_id, 
match_date, round_name, lichess_round_id, lichess_game_id, match_result) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

