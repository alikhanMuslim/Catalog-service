-- name: CreateUser :one
INSERT INTO users (
  username, password_hash, email, token
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;