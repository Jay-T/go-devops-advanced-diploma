// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: files.sql

package db

import (
	"context"
)

const createFile = `-- name: CreateFile :one
INSERT INTO files (
  account_id,
  filename,
  filepath,
  filesize
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, account_id, filename, filepath, filesize, deleted, created_at
`

type CreateFileParams struct {
	AccountID int64  `json:"account_id"`
	Filename  string `json:"filename"`
	Filepath  string `json:"filepath"`
	Filesize  int64  `json:"filesize"`
}

func (q *Queries) CreateFile(ctx context.Context, arg CreateFileParams) (File, error) {
	row := q.db.QueryRowContext(ctx, createFile,
		arg.AccountID,
		arg.Filename,
		arg.Filepath,
		arg.Filesize,
	)
	var i File
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Filename,
		&i.Filepath,
		&i.Filesize,
		&i.Deleted,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFile = `-- name: DeleteFile :exec
UPDATE files
  set deleted = true
WHERE filename = $1 and filepath = $2 and account_id = $3
`

type DeleteFileParams struct {
	Filename  string `json:"filename"`
	Filepath  string `json:"filepath"`
	AccountID int64  `json:"account_id"`
}

func (q *Queries) DeleteFile(ctx context.Context, arg DeleteFileParams) error {
	_, err := q.db.ExecContext(ctx, deleteFile, arg.Filename, arg.Filepath, arg.AccountID)
	return err
}

const deletedFileById = `-- name: DeletedFileById :exec
DELETE FROM files
WHERE id = $1
`

func (q *Queries) DeletedFileById(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletedFileById, id)
	return err
}

const getDeletedFiles = `-- name: GetDeletedFiles :many
SELECT f.id, f.filename, f.filepath, a.username 
FROM files f JOIN account a ON f.account_id = a.id
WHERE f.deleted = true
`

type GetDeletedFilesRow struct {
	ID       int64  `json:"id"`
	Filename string `json:"filename"`
	Filepath string `json:"filepath"`
	Username string `json:"username"`
}

func (q *Queries) GetDeletedFiles(ctx context.Context) ([]GetDeletedFilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getDeletedFiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDeletedFilesRow
	for rows.Next() {
		var i GetDeletedFilesRow
		if err := rows.Scan(
			&i.ID,
			&i.Filename,
			&i.Filepath,
			&i.Username,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFile = `-- name: GetFile :one
SELECT id, account_id, filename, filepath, filesize, deleted, created_at FROM files
WHERE filename = $1 and filepath = $2 and account_id = $3 and deleted = false LIMIT 1
`

type GetFileParams struct {
	Filename  string `json:"filename"`
	Filepath  string `json:"filepath"`
	AccountID int64  `json:"account_id"`
}

func (q *Queries) GetFile(ctx context.Context, arg GetFileParams) (File, error) {
	row := q.db.QueryRowContext(ctx, getFile, arg.Filename, arg.Filepath, arg.AccountID)
	var i File
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Filename,
		&i.Filepath,
		&i.Filesize,
		&i.Deleted,
		&i.CreatedAt,
	)
	return i, err
}

const listFiles = `-- name: ListFiles :many
SELECT id, account_id, filename, filepath, filesize, deleted, created_at FROM files
WHERE account_id = $1 and deleted = false 
ORDER BY filename
`

func (q *Queries) ListFiles(ctx context.Context, accountID int64) ([]File, error) {
	rows, err := q.db.QueryContext(ctx, listFiles, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []File
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Filename,
			&i.Filepath,
			&i.Filesize,
			&i.Deleted,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFileName = `-- name: UpdateFileName :exec
UPDATE files
  set filename = $4
WHERE filename = $1 and filepath = $2 and account_id = $3 and deleted = false
`

type UpdateFileNameParams struct {
	Filename   string `json:"filename"`
	Filepath   string `json:"filepath"`
	AccountID  int64  `json:"account_id"`
	Filename_2 string `json:"filename_2"`
}

func (q *Queries) UpdateFileName(ctx context.Context, arg UpdateFileNameParams) error {
	_, err := q.db.ExecContext(ctx, updateFileName,
		arg.Filename,
		arg.Filepath,
		arg.AccountID,
		arg.Filename_2,
	)
	return err
}
