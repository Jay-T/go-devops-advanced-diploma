// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: secrets_metadata.sql

package db

import (
	"context"
)

const createSecretMetadata = `-- name: CreateSecretMetadata :one
INSERT INTO secrets_metadata (
  secret_id,
  key,
  value
) VALUES (
  $1, $2, $3
)
RETURNING id, secret_id, key, value, created_at
`

type CreateSecretMetadataParams struct {
	SecretID int64  `json:"secret_id"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

func (q *Queries) CreateSecretMetadata(ctx context.Context, arg CreateSecretMetadataParams) (SecretsMetadatum, error) {
	row := q.db.QueryRowContext(ctx, createSecretMetadata, arg.SecretID, arg.Key, arg.Value)
	var i SecretsMetadatum
	err := row.Scan(
		&i.ID,
		&i.SecretID,
		&i.Key,
		&i.Value,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSecretMetadata = `-- name: DeleteSecretMetadata :exec
DELETE FROM secrets_metadata
WHERE key = $1 and secret_id = $2
`

type DeleteSecretMetadataParams struct {
	Key      string `json:"key"`
	SecretID int64  `json:"secret_id"`
}

func (q *Queries) DeleteSecretMetadata(ctx context.Context, arg DeleteSecretMetadataParams) error {
	_, err := q.db.ExecContext(ctx, deleteSecretMetadata, arg.Key, arg.SecretID)
	return err
}

const listSecretMetadata = `-- name: ListSecretMetadata :many
SELECT id, secret_id, key, value, created_at FROM secrets_metadata
WHERE secret_id = $1 
ORDER BY key
`

func (q *Queries) ListSecretMetadata(ctx context.Context, secretID int64) ([]SecretsMetadatum, error) {
	rows, err := q.db.QueryContext(ctx, listSecretMetadata, secretID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SecretsMetadatum
	for rows.Next() {
		var i SecretsMetadatum
		if err := rows.Scan(
			&i.ID,
			&i.SecretID,
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
UPDATE secrets_metadata
  set value = $3
WHERE key = $1 and secret_id = $2
`

type UpdateSecretMetadataParams struct {
	Key      string `json:"key"`
	SecretID int64  `json:"secret_id"`
	Value    string `json:"value"`
}

func (q *Queries) UpdateSecretMetadata(ctx context.Context, arg UpdateSecretMetadataParams) error {
	_, err := q.db.ExecContext(ctx, updateSecretMetadata, arg.Key, arg.SecretID, arg.Value)
	return err
}
