// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: metadata.sql

package db

import (
	"context"
)

const createSecretMetadata = `-- name: CreateSecretMetadata :one
INSERT INTO metadata (
  obj_id,
  obj_type,
  account_id,
  key,
  value
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, obj_id, obj_type, account_id, key, value, created_at
`

type CreateSecretMetadataParams struct {
	ObjID     int64  `json:"obj_id"`
	ObjType   string `json:"obj_type"`
	AccountID int64  `json:"account_id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func (q *Queries) CreateSecretMetadata(ctx context.Context, arg CreateSecretMetadataParams) (Metadatum, error) {
	row := q.db.QueryRowContext(ctx, createSecretMetadata,
		arg.ObjID,
		arg.ObjType,
		arg.AccountID,
		arg.Key,
		arg.Value,
	)
	var i Metadatum
	err := row.Scan(
		&i.ID,
		&i.ObjID,
		&i.ObjType,
		&i.AccountID,
		&i.Key,
		&i.Value,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAllSecretMetadata = `-- name: DeleteAllSecretMetadata :exec
DELETE FROM metadata
WHERE obj_id = $1 and account_id = $2 and obj_type = $3
`

type DeleteAllSecretMetadataParams struct {
	ObjID     int64  `json:"obj_id"`
	AccountID int64  `json:"account_id"`
	ObjType   string `json:"obj_type"`
}

func (q *Queries) DeleteAllSecretMetadata(ctx context.Context, arg DeleteAllSecretMetadataParams) error {
	_, err := q.db.ExecContext(ctx, deleteAllSecretMetadata, arg.ObjID, arg.AccountID, arg.ObjType)
	return err
}

const deleteOneSecretMetadata = `-- name: DeleteOneSecretMetadata :exec
DELETE FROM metadata
WHERE obj_id = $1 and account_id = $2 and obj_type = $3 and key = $4
`

type DeleteOneSecretMetadataParams struct {
	ObjID     int64  `json:"obj_id"`
	AccountID int64  `json:"account_id"`
	ObjType   string `json:"obj_type"`
	Key       string `json:"key"`
}

func (q *Queries) DeleteOneSecretMetadata(ctx context.Context, arg DeleteOneSecretMetadataParams) error {
	_, err := q.db.ExecContext(ctx, deleteOneSecretMetadata,
		arg.ObjID,
		arg.AccountID,
		arg.ObjType,
		arg.Key,
	)
	return err
}

const listSecretMetadata = `-- name: ListSecretMetadata :many
SELECT id, obj_id, obj_type, account_id, key, value, created_at FROM metadata
WHERE obj_id = $1 and account_id = $2 and obj_type = $3
ORDER BY key
`

type ListSecretMetadataParams struct {
	ObjID     int64  `json:"obj_id"`
	AccountID int64  `json:"account_id"`
	ObjType   string `json:"obj_type"`
}

func (q *Queries) ListSecretMetadata(ctx context.Context, arg ListSecretMetadataParams) ([]Metadatum, error) {
	rows, err := q.db.QueryContext(ctx, listSecretMetadata, arg.ObjID, arg.AccountID, arg.ObjType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Metadatum
	for rows.Next() {
		var i Metadatum
		if err := rows.Scan(
			&i.ID,
			&i.ObjID,
			&i.ObjType,
			&i.AccountID,
			&i.Key,
			&i.Value,
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

const updateSecretMetadata = `-- name: UpdateSecretMetadata :exec
UPDATE metadata
  set value = $5
WHERE obj_id = $1 and account_id = $2 and obj_type = $3 and key = $4
`

type UpdateSecretMetadataParams struct {
	ObjID     int64  `json:"obj_id"`
	AccountID int64  `json:"account_id"`
	ObjType   string `json:"obj_type"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func (q *Queries) UpdateSecretMetadata(ctx context.Context, arg UpdateSecretMetadataParams) error {
	_, err := q.db.ExecContext(ctx, updateSecretMetadata,
		arg.ObjID,
		arg.AccountID,
		arg.ObjType,
		arg.Key,
		arg.Value,
	)
	return err
}
