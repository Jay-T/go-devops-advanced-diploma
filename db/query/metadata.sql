-- name: CreateSecretMetadata :one
INSERT INTO metadata (
  obj_id,
  obj_type,
  account_id,
  key,
  value
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateSecretMetadata :exec
UPDATE metadata
  set value = $5
WHERE obj_id = $1 and account_id = $2 and obj_type = $3 and key = $4;

-- name: ListSecretMetadata :many
SELECT * FROM metadata
WHERE obj_id = $1 and account_id = $2 and obj_type = $3
ORDER BY key;

-- name: DeleteOneSecretMetadata :exec
DELETE FROM metadata
WHERE obj_id = $1 and account_id = $2 and obj_type = $3 and key = $4;

-- name: DeleteAllSecretMetadata :exec
DELETE FROM metadata
WHERE obj_id = $1 and account_id = $2 and obj_type = $3;