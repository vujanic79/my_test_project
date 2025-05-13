-- name: CreateUser :one
INSERT INTO app_user (id, created_at, updated_at, first_name, last_name, email)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserIdByEmail :one
SELECT u.id FROM app_user u
WHERE u.email = $1;