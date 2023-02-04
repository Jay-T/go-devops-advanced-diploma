package db

import (
	"context"

	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (store *SQLStore) CreateFileTx(ctx context.Context, arg CreateFileParams, errChan <-chan error) (File, error) {
	var file File

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		file, err = store.CreateFile(ctx, arg)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "unique_violation":
					return status.Errorf(codes.AlreadyExists, "file already exists: %s", err)
				}
			}
			return status.Errorf(codes.Internal, "failed to create file: Err: %s", err)
		}

		err = <-errChan

		return err
	})

	return file, err
}
