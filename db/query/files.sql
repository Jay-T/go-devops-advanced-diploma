-- name: CreateFile :one
INSERT INTO files (
  account_id,
  filename,
  filepath,
  ready
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateFilePath :exec
UPDATE files
  set filepath = $3
WHERE filename = $1 and account_id = $2;

-- name: MarkFileReady :exec
UPDATE files
  set ready = true
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