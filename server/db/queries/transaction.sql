-- name: GetTransactionById :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: CreateTransaction :one
INSERT INTO transactions (
    ledger, debitor, creditor, amount, note, date
)
VALUES (
    $1, $2, $3, $4, $5, $6
)
ON CONFLICT DO NOTHING
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;