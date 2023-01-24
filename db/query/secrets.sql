-- name: CreateSecret :one
INSERT INTO secrets (
  account_id,
  key,
  value
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateSecret :exec
UPDATE secrets
  set value = $3
WHERE key = $1 and account_id = $2;

-- name: GetSecret :one
SELECT * FROM secrets
WHERE key = $1 and account_id = $2 LIMIT 1;

-- name: ListSecrets :many
SELECT * FROM secrets
WHERE account_id = $1 
ORDER BY key;

-- name: DeleteSecret :exec
DELETE FROM secrets
WHERE key = $1 and account_id = $2;