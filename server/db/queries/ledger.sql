-- name: GetLedgerById :one
SELECT ledger
FROM (
  SELECT l.id, l.title, json_agg(json_build_object(t.*)) AS transactions
  FROM ledgers l
    JOIN transactions t ON l.id = t.ledger
  WHERE l.id = $1
  GROUP BY l.id
) ledger;

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
UPDATE ledgers SET members = ARRAY_APPEND(members, $2)
WHERE id = $1;

-- name: RemoveUserFromLedger :exec
UPDATE ledgers SET members = ARRAY_REMOVE(members, $2)
WHERE ledgers.id::text = $1;

-- name: GetTransactions :many
SELECT * FROM transactions
WHERE transactions.ledger = $1;