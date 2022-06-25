-- name: GetLedgerById :one
SELECT * FROM ledgers
WHERE id = $1 LIMIT 1;

-- name: CreateLedger :one
INSERT INTO ledgers (
    title, members
)
VALUES (
    $1, $2
)
ON CONFLICT DO NOTHING
RETURNING *;

-- name: DeleteLedger :exec
DELETE FROM ledgers
WHERE id = $1;

-- name: AddUserToLedger :exec
UPDATE ledgers SET members = array_append(members, $2)
WHERE id = $1
RETURNING *;

-- name: AddTransactionToLedger :exec
UPDATE ledgers SET transactions = array_append(transactions, $2)
WHERE id = $1
RETURNING *;

-- name: RemoveUserFromLedger :exec
SELECT array_remove(members, $2) as members from ledgers
WHERE id = $1;