-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- -- name: GetUserLedgers :many
-- SELECT * FROM ledgers
-- WHERE $1 IN members;

-- name: CreateUser :one
INSERT INTO users (
    email, first_name, last_name
)
VALUES (
    $1, $2, $3
)
ON CONFLICT DO NOTHING
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;