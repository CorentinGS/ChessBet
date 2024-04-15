-- name: GetTournament :one
SELECT * FROM tournaments WHERE tournament_id = $1;

-- name: GetTournaments :many
SELECT * FROM tournaments;

-- name: CreateTournament :one
INSERT INTO tournaments (name, start_date, end_date) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateTournament :one
UPDATE tournaments SET name = $2, start_date = $3, end_date = $4 WHERE tournament_id = $1 RETURNING *;

-- name: DeleteTournament :one
DELETE FROM tournaments WHERE tournament_id = $1 RETURNING *;

-- name: GetTournamentInProgress :many
SELECT * FROM tournaments WHERE start_date <= CURRENT_DATE AND end_date >= CURRENT_DATE;

-- name: GetTournamentUpcoming :many
SELECT * FROM tournaments WHERE start_date > CURRENT_DATE;

-- name: GetTournamentPast :many
SELECT * FROM tournaments WHERE end_date < CURRENT_DATE;
