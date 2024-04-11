-- name: GetUser :one
SELECT * FROM users WHERE user_id = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (username, points, email_address) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET username = $2, points = $3 WHERE user_id = $1 RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE user_id = $1 RETURNING *;

-- name: GetUserByOauthId :one
SELECT * FROM users WHERE oauth_id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email_address = $1;