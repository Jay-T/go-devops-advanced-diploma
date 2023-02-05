// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	BlockAccount(ctx context.Context, username string) error
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateFile(ctx context.Context, arg CreateFileParams) (File, error)
	CreateOrUpdateFileMetadata(ctx context.Context, arg CreateOrUpdateFileMetadataParams) (Metadatum, error)
	CreateOrUpdateSecretMetadata(ctx context.Context, arg CreateOrUpdateSecretMetadataParams) (Metadatum, error)
	CreateSecret(ctx context.Context, arg CreateSecretParams) (Secret, error)
	DeleteAccount(ctx context.Context, username string) error
	DeleteAllFileMetadata(ctx context.Context, fileID sql.NullInt64) error
	DeleteAllSecretMetadata(ctx context.Context, secretID sql.NullInt64) error
	DeleteFile(ctx context.Context, arg DeleteFileParams) error
	DeleteOneFileMetadata(ctx context.Context, arg DeleteOneFileMetadataParams) error
	DeleteOneSecretMetadata(ctx context.Context, arg DeleteOneSecretMetadataParams) error
	DeleteSecret(ctx context.Context, arg DeleteSecretParams) error
	DeletedFileById(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, username string) (Account, error)
	GetDeletedFiles(ctx context.Context) ([]GetDeletedFilesRow, error)
	GetFile(ctx context.Context, arg GetFileParams) (File, error)
	GetSecret(ctx context.Context, arg GetSecretParams) (Secret, error)
	ListFileMetadata(ctx context.Context, fileID sql.NullInt64) ([]Metadatum, error)
	ListFiles(ctx context.Context, accountID int64) ([]File, error)
	ListSecretMetadata(ctx context.Context, secretID sql.NullInt64) ([]Metadatum, error)
	ListSecrets(ctx context.Context, accountID int64) ([]Secret, error)
	UpdateFileName(ctx context.Context, arg UpdateFileNameParams) (File, error)
	UpdateSecret(ctx context.Context, arg UpdateSecretParams) (Secret, error)
}

var _ Querier = (*Queries)(nil)
