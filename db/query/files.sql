-- name: CreateFile :one
INSERT INTO files (
  account_id,
  filename,
  filepath,
  filesize
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateFileName :one
UPDATE files
  set filename = $4
WHERE filename = $1 and filepath = $2 and account_id = $3 and deleted = false
RETURNING *;

-- name: GetFile :one
SELECT * FROM files
WHERE filename = $1 and filepath = $2 and account_id = $3 and deleted = false LIMIT 1;

-- name: ListFiles :many
SELECT * FROM files
WHERE account_id = $1 and deleted = false 
ORDER BY filename;

-- name: DeleteFile :exec
UPDATE files
  set deleted = true
WHERE filename = $1 and filepath = $2 and account_id = $3;

-- name: DeletedFileById :exec
DELETE FROM files
WHERE id = $1;

-- name: GetDeletedFiles :many
SELECT f.id, f.filename, f.filepath, a.username 
FROM files f JOIN account a ON f.account_id = a.id
WHERE f.deleted = true;
