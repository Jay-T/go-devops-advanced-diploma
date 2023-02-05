-- name: CreateOrUpdateSecretMetadata :one
INSERT INTO metadata (
  secret_id,
  key,
  value
) VALUES (
  $1, $2, $3
)
ON CONFLICT(secret_id, key) 
DO UPDATE
 set value = $3
RETURNING *;

-- name: CreateOrUpdateFileMetadata :one
INSERT INTO metadata (
  file_id,
  key,
  value
) VALUES (
  $1, $2, $3
)
ON CONFLICT(file_id, key) 
DO UPDATE
 set value = $3
RETURNING *;

-- name: ListSecretMetadata :many
SELECT * FROM metadata
WHERE secret_id = $1 
ORDER BY key;

-- name: ListFileMetadata :many
SELECT * FROM metadata
WHERE file_id = $1
ORDER BY key;

-- name: DeleteOneSecretMetadata :exec
DELETE FROM metadata
WHERE secret_id = $1 and key = $2;

-- name: DeleteOneFileMetadata :exec
DELETE FROM metadata
WHERE file_id = $1 and key = $2;

-- name: DeleteAllSecretMetadata :exec
DELETE FROM metadata
WHERE secret_id = $1;

-- name: DeleteAllFileMetadata :exec
DELETE FROM metadata
WHERE file_id = $1;