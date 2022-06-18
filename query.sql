-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users LIMIT 100;
