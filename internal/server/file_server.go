package server

import (
	"bytes"
	"context"
	"io"

	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	maxImageSize = 1 << 10
)

type FileServer struct {
	fileStore        db.Store
	fileContentSaver FileContentSaver
	pb.UnimplementedFileServer
}

func NewFileServer(fileStore db.Store) *FileServer {
	fileContentSaver := NewDiskFileContentSaver("fs")
	return &FileServer{
		fileStore,
		fileContentSaver,
		pb.UnimplementedFileServer{},
	}
}

// // findAccount gets account info from Accounts table.
// func (s *AuthServer) findAccount(ctx context.Context) (db.Account, error) {
// 	username, err := getUsernameFromContext(ctx)
// 	if err != nil {
// 		return db.Account{}, err
// 	}

// 	account, err := s.accountStore.GetAccount(ctx, username)
// 	if err != nil {
// 		return db.Account{}, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
// 	}

// 	return account, nil
// }

func (s *FileServer) CreateFile(stream pb.File_CreateFileServer) error {
	// TODO(): make as TX !!!
	ctx := stream.Context()
	account, err := findAccount(ctx, s.fileStore)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("receive an CreateFile request from user %s", account.Username)

	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive file info"))
	}
	arg := db.CreateFileParams{
		AccountID: account.ID,
		Filename:  req.GetInfo().Filename,
		Filepath:  req.GetInfo().Filepath,
	}

	var file db.File
	errChan := make(chan error)
	go func() {
		file, err = s.fileStore.CreateFileTx(ctx, arg, errChan)
	}()
	// if err != nil {
	// 	if pqErr, ok := err.(*pq.Error); ok {
	// 		switch pqErr.Code.Name() {
	// 		case "unique_violation":
	// 			return logError(status.Errorf(codes.AlreadyExists, "file already exists: %s", err))
	// 		}
	// 	}
	// 	return logError(status.Errorf(codes.Internal, "failed to create file: Err: %s", err))
	// }

	fileData := bytes.Buffer{}
	fileSize := 0

	err = func() error {
		for {
			err := contextError(stream.Context())
			if err != nil {
				return err
			}
			log.Info().Msg("waiting to receive more filedata")

			req, err := stream.Recv()
			if err == io.EOF {
				log.Print("no more data")
				break
			}
			if err != nil {
				return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
			}

			chunk := req.GetChunkData()
			size := len(chunk)

			log.Printf("received a chunk with size: %d", size)

			fileSize += size
			if fileSize > maxImageSize {
				return logError(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", fileSize, maxImageSize))
			}

			_, err = fileData.Write(chunk)
			if err != nil {
				return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
			}
		}
		return nil
	}()
	if err != nil {
		errChan <- err
		return err
	}
	// if err != nil {
	// 	argDelete := db.DeleteFileParams{
	// 		Filename:  arg.Filename,
	// 		AccountID: arg.AccountID,
	// 	}
	// 	err = s.fileStore.DeleteFile(ctx, argDelete)
	// 	if err != nil {
	// 		return logError(status.Errorf(codes.Internal, "cannot delete file record from db: %v", err))
	// 	}
	// 	return err
	// }

	err = s.fileContentSaver.Save(req.GetInfo().Filename, req.GetInfo().Filepath, fileData)
	if err != nil {
		errChan <- err
		return logError(status.Errorf(codes.Internal, "cannot save file content to storage: %v", err))
	}

	errChan <- nil

	res := &pb.CreateFileResponse{
		Info: &pb.FileInfo{
			Filename: req.GetInfo().Filename,
			Filepath: req.GetInfo().Filepath,
		},
		Size: uint32(fileSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	log.Info().Msgf("Created file '/%s/%s' with size %s", file.Filename, arg.Filepath, fileSize)
	return nil
}

func (s *FileServer) UpdateFile(stream pb.File_UpdateFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}

func (s *FileServer) DeleteFile(context.Context, *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}

func (s *FileServer) GetFile(*pb.GetFileRequest, pb.File_GetFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}

func (s *FileServer) ListFile(context.Context, *pb.ListFileRequest) (*pb.ListFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFile not implemented")
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

// logError returns formatted error.
func logError(err error) error {
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}
