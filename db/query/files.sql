-- name: CreateFile :one
INSERT INTO files (
  account_id,
  filename,
  filepath
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateFilePath :exec
UPDATE files
  set filepath = $3
WHERE filename = $1 and account_id = $2;

-- name: GetFile :one
SELECT * FROM files
WHERE filename = $1 and account_id = $2 LIMIT 1;

-- name: ListFiles :many
SELECT * FROM files
WHERE account_id = $1 
ORDER BY filename;

-- name: DeleteFile :exec
DELETE FROM files
WHERE filename = $1 and account_id = $2;