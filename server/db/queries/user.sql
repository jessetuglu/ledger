-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserLedgers :many
SELECT * FROM ledgers
WHERE $1 = ANY (ledgers.members);

-- name: GetOrCreateUser :one
WITH i AS(
    INSERT INTO users (email, first_name, last_name) 
    VALUES ($1, $2, $3)
    ON CONFLICT(email) DO NOTHING
    RETURNING *
)
SELECT * FROM i
UNION
SELECT * FROM users WHERE email = $1;