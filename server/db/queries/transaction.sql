-- name: GetTransactionById :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: CreateTransaction :one
INSERT INTO transactions (
    debitor, creditor, amount, note
)
VALUES (
    $1, $2, $3, $4
)
ON CONFLICT DO NOTHING
RETURNING *;