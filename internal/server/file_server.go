package server

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"time"

	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
	"github.com/Jay-T/go-devops-advanced-diploma/internal/pb"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	maxImageSize = 1 << 10
)

type FileServer struct {
	fileStore        db.Store
	fileContentSaver FileContentSaver
	pb.UnimplementedFileServer
}

func NewFileServer(ctx context.Context, fileStore db.Store) *FileServer {
	fileContentSaver := NewDiskFileContentSaver("fs")
	s := &FileServer{
		fileStore,
		fileContentSaver,
		pb.UnimplementedFileServer{},
	}
	go s.ClearFileStorage(ctx)

	return s
}

// CreateFile creates file in storage for user.
func (s *FileServer) CreateFile(stream pb.File_CreateFileServer) error {
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
	filename := req.GetInfo().Filename
	filepath := req.GetInfo().Filepath

	arg := db.GetFileParams{
		AccountID: account.ID,
		Filename:  filename,
		Filepath:  filepath,
	}

	file, err := s.fileStore.GetFile(ctx, arg)
	if err != nil && err != sql.ErrNoRows {
		return logError(status.Errorf(codes.Internal, "failed to find the file: Err: %s", err))
	}
	if file.Filename != "" {
		return logError(status.Errorf(codes.AlreadyExists, "file already exists"))
	}

	fileData := bytes.Buffer{}
	fileSize := 0

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

	fsFilepath := fmt.Sprintf("%s/%s", account.Username, filepath)
	err = s.fileContentSaver.Save(ctx, filename, fsFilepath, fileData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save file content to storage: %v", err))
	}

	argCreateFile := db.CreateFileParams{
		AccountID: account.ID,
		Filename:  filename,
		Filepath:  filepath,
		Filesize:  int64(fileSize),
	}

	newFile, err := s.fileStore.CreateFile(ctx, argCreateFile)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return logError(status.Errorf(codes.AlreadyExists, "Secret already exists: %s", err))
			}
		}
		return logError(status.Errorf(codes.Internal, "failed to create secret: Err: %s", err))
	}

	res := &pb.CreateFileResponse{
		Info: &pb.FileInfo{
			Filename: newFile.Filename,
			Filepath: newFile.Filepath,
			Size:     toUint64Ref(newFile.Filesize),
		},
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	log.Info().Msgf("Created file '/%s/%s' with size %s", filename, filepath, fileSize)
	return nil
}

// DeleteFile deletes file from storage.
func (s *FileServer) DeleteFile(ctx context.Context, in *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	account, err := findAccount(ctx, s.fileStore)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("receive an DeleteFile request from user %s", account.Username)

	arg := db.DeleteFileParams{
		Filename:  in.Info.Filename,
		Filepath:  in.Info.Filepath,
		AccountID: account.ID,
	}

	err = s.fileStore.DeleteFile(ctx, arg)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot delete file from db. Err :%s", err))
	}

	resp := &pb.DeleteFileResponse{
		Info: &pb.FileInfo{
			Filename: in.Info.Filename,
			Filepath: in.Info.Filepath,
		},
	}
	return resp, nil
}

// UpdateFileName allows to change file name.
func (s *FileServer) UpdateFileName(ctx context.Context, in *pb.UpdateFileNameRequest) (*pb.UpdateFileNameResponse, error) {
	account, err := findAccount(ctx, s.fileStore)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("receive an UpdateFileName request from user %s", account.Username)

	oldName := in.Info.Filename
	newName := in.NewFilename
	path := in.Info.Filepath

	argGetFile := db.GetFileParams{
		AccountID: account.ID,
		Filename:  oldName,
		Filepath:  path,
	}

	file, err := s.fileStore.GetFile(ctx, argGetFile)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "failed to find the file: Err: %s", err))
	}

	pathFS := fmt.Sprintf("%s/%s", account.Username, path)
	err = s.fileContentSaver.Rename(ctx, oldName, newName, pathFS)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot rename file in fs. Err :%s", err))
	}

	arg := db.UpdateFileNameParams{
		Filename:   oldName,
		Filepath:   path,
		AccountID:  account.ID,
		Filename_2: newName,
	}

	err = s.fileStore.UpdateFileName(ctx, arg)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot rename file in db. Err :%s", err))
	}

	resp := &pb.UpdateFileNameResponse{
		Info: &pb.FileInfo{
			Filename:  newName,
			Filepath:  in.Info.Filepath,
			Size:      toUint64Ref(file.Filesize),
			CreatedAt: timestamppb.New(file.CreatedAt),
		},
	}

	return resp, nil
}

// GetFile returns a file.
func (s *FileServer) GetFile(*pb.GetFileRequest, pb.File_GetFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}

// ListFiles lists all files belong to user.
func (s *FileServer) ListFiles(context.Context, *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFile not implemented")
}

// contextError handles context events.
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

func (s *FileServer) ClearFileStorage(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			log.Info().Msg("Running garbage collector for deleted files.")

			deletedFiles, err := s.fileStore.GetDeletedFiles(ctx)
			if err != nil {
				log.Error().Msgf("could not get deleted files from db. Err: %s", err)
			}

			for _, file := range deletedFiles {
				filepath := fmt.Sprintf("%s/%s", file.Username, file.Filepath)
				err := s.fileContentSaver.Delete(ctx, file.Filename, filepath)
				if err != nil {
					log.Error().Msgf("could not delete file '%s' at path '%s'", file.Filename, filepath)
					continue
				}

				err = s.fileStore.DeletedFileById(ctx, file.ID)
				if err != nil {
					log.Error().Msgf("could not delete file with ID '%d' from db", file.ID)
					continue
				}

				log.Info().Msgf("Deleted file '%s' at path '%s'", file.Filename, filepath)
			}

		case <-ctx.Done():
			log.Info().Msg("ClearFileStorage has been canceled successfully.")
			return
		}
	}
}

// logError returns formatted error.
func logError(err error) error {
	if err != nil {
		log.Error().Err(err).Msg("")
	}
	return err
}

func toUint64Ref(i int64) *uint64 {
	ii := uint64(i)
	return &ii
}
