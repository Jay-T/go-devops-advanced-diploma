-- name: CreateSecretMetadata :one
INSERT INTO secrets_metadata (
  secret_id,
  key,
  value
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateSecretMetadata :exec
UPDATE secrets_metadata
  set value = $3
WHERE key = $1 and secret_id = $2;

-- name: ListSecretMetadata :many
SELECT * FROM secrets_metadata
WHERE secret_id = $1 
ORDER BY key;

-- name: DeleteSecretMetadata :exec
DELETE FROM secrets_metadata
WHERE key = $1 and secret_id = $2;