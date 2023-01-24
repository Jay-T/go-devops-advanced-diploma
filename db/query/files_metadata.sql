-- name: CreateFileMetadata :one
INSERT INTO files_metadata (
  file_id,
  key,
  value
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateFileMetadata :exec
UPDATE files_metadata
  set value = $3
WHERE key = $1 and file_id = $2;

-- name: ListFileMetadata :many
SELECT * FROM files_metadata
WHERE file_id = $1 
ORDER BY key;

-- name: DeleteFileMetadata :exec
DELETE FROM files_metadata
WHERE key = $1 and file_id = $2;