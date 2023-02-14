package server

import (
	"bufio"
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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	maxImageSize = 1 << 10
)

// FileServer struct describes FileServer fields.
type FileServer struct {
	fileStore        db.Store
	fileContentSaver FileContentSaver
	pb.UnimplementedFileServer
}

// NewFileServer returns a FileServer instance.
func NewFileServer(ctx context.Context, fileStore db.Store, FSroot string) *FileServer {
	fileContentSaver := NewDiskFileContentSaver(FSroot)
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
	metadataList := req.GetInfo().Metadata

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
		log.Debug().Msg("waiting to receive more filedata")

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

	if len(metadataList) != 0 {
		for _, md := range metadataList {
			argMD := db.CreateOrUpdateFileMetadataParams{
				FileID: SQLInt64(newFile.ID),
				Key:    md.Key,
				Value:  md.Value,
			}
			_, err := s.fileStore.CreateOrUpdateFileMetadata(ctx, argMD)
			if err != nil {
				return logError(status.Errorf(codes.Internal, "failed to create file metadata: Err: %s", err))
			}
		}
	}

	res := &pb.CreateFileResponse{
		Info: &pb.FileInfo{
			Filename:  newFile.Filename,
			Filepath:  newFile.Filepath,
			Size:      toUint64Ref(newFile.Filesize),
			Metadata:  metadataList,
			CreatedAt: timestamppb.New(newFile.CreatedAt),
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

	_, err = s.fileStore.GetFile(ctx, argGetFile)
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

	fileAfterUpdate, err := s.fileStore.UpdateFileName(ctx, arg)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot rename file in db. Err :%s", err))
	}

	metadataList := in.Info.GetMetadata()
	newMDList, err := s.CreateOrUpdateFileMD(ctx, metadataList, SQLInt64(fileAfterUpdate.ID))
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateFileNameResponse{
		Info: &pb.FileInfo{
			Filename:  fileAfterUpdate.Filename,
			Filepath:  fileAfterUpdate.Filepath,
			Size:      toUint64Ref(fileAfterUpdate.Filesize),
			CreatedAt: timestamppb.New(fileAfterUpdate.CreatedAt),
			Metadata:  ConvertToPBMetadata(newMDList),
		},
	}

	return resp, nil
}

// ListFiles lists all files belong to user.
func (s *FileServer) ListFiles(ctx context.Context, in *emptypb.Empty) (*pb.ListFilesResponse, error) {
	account, err := findAccount(ctx, s.fileStore)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("receive an UpdateFileName request from user %s", account.Username)

	files, err := s.fileStore.ListFiles(ctx, account.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, logError(status.Errorf(codes.NotFound, "no files found."))
		}
		return nil, logError(status.Errorf(codes.Internal, "cannot get files. Err :%s", err))
	}
	var pbFiles []*pb.FileInfo
	for _, file := range files {
		metadata, err := s.fileStore.ListFileMetadata(ctx, SQLInt64(file.ID))
		if err != nil && err != sql.ErrNoRows {
			return nil, logError(status.Errorf(codes.Internal, "cannot collect secret metadata. Err: %s", err))
		}

		pbFiles = append(pbFiles, &pb.FileInfo{
			Filepath:  file.Filepath,
			Filename:  file.Filename,
			Size:      toUint64Ref(file.Filesize),
			CreatedAt: timestamppb.New(file.CreatedAt),
			Metadata:  ConvertToPBMetadata(metadata),
		})
	}

	return &pb.ListFilesResponse{
		Info: pbFiles,
	}, nil
}

// GetFileInfo returns file info.
func (s *FileServer) GetFileInfo(ctx context.Context, in *pb.GetFileInfoRequest) (*pb.GetFileInfoResponse, error) {
	account, err := findAccount(ctx, s.fileStore)
	if err != nil {
		return nil, logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("receive an GetFileInfo request from user %s", account.Username)

	args := db.GetFileParams{
		Filename:  in.Info.GetFilename(),
		Filepath:  in.Info.GetFilepath(),
		AccountID: account.ID,
	}

	file, err := s.fileStore.GetFile(ctx, args)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, logError(status.Errorf(codes.NotFound, "file not found."))
		}
		return nil, logError(status.Errorf(codes.Internal, "cannot get file info from db. Err :%s", err))
	}

	metadata, err := s.fileStore.ListFileMetadata(ctx, SQLInt64(file.ID))
	if err != nil && err != sql.ErrNoRows {
		return nil, logError(status.Errorf(codes.Internal, "cannot collect file metadata. Err: %s", err))
	}

	fileInfo := &pb.FileInfo{
		Filename:  file.Filename,
		Filepath:  file.Filepath,
		Size:      toUint64Ref(file.Filesize),
		CreatedAt: timestamppb.New(file.CreatedAt),
		Metadata:  ConvertToPBMetadata(metadata),
	}

	resp := &pb.GetFileInfoResponse{
		Info: fileInfo,
	}

	return resp, nil
}

// DownloadFile returns requested file content.
func (s *FileServer) DownloadFile(in *pb.DownloadFileRequest, stream pb.File_DownloadFileServer) error {
	ctx := stream.Context()
	account, err := findAccount(ctx, s.fileStore)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot get account from db. Err :%s", err))
	}

	log.Info().Msgf("receive an DownloadFile request from user %s", account.Username)

	filename := in.Info.GetFilename()
	filepath := in.Info.GetFilepath()

	argFileInfo := db.GetFileParams{
		Filename:  filename,
		Filepath:  filepath,
		AccountID: account.ID,
	}
	fileInfo, err := s.fileStore.GetFile(ctx, argFileInfo)
	if err != nil {
		if err == sql.ErrNoRows {
			return logError(status.Errorf(codes.NotFound, "file not found."))
		}
		return logError(status.Errorf(codes.Internal, "cannot get file info from db. Err :%s", err))
	}

	pathFS := fmt.Sprintf("%s/%s", account.Username, fileInfo.Filepath)
	file, err := s.fileContentSaver.LoadFile(ctx, fileInfo.Filename, pathFS)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot read file. Err :%s", err))
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Internal, "error reading file content. Err :%s", err))
		}

		resp := &pb.DownloadFileResponse{
			ChunkData: buffer[:n],
			BytesSent: int32(len(buffer[:n])),
		}

		err = stream.Send(resp)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot send chunk to client: %s, %s", err, stream.RecvMsg(nil)))
		}
	}

	log.Info().Msgf("finished sending file '%s' to client", fileInfo.Filename)

	return nil
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

// ClearFileStorage deletes files from storage which marked as deleted=true in DB.
func (s *FileServer) ClearFileStorage(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			log.Debug().Msg("Running garbage collector for deleted files.")

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

				err = s.fileStore.DeleteAllFileMetadata(ctx, SQLInt64(file.ID))
				if err != nil {
					logError(status.Errorf(codes.Internal, "cannot delete file metadata: Err: %s", err))
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

// CreateOrUpdateFileMD updates metadata for file.
func (s *FileServer) CreateOrUpdateFileMD(ctx context.Context, metadataList []*pb.Metadata, fileID sql.NullInt64) ([]db.Metadatum, error) {
	newMDList := []db.Metadatum{}
	if len(metadataList) != 0 {
		for _, md := range metadataList {
			argMD := db.CreateOrUpdateFileMetadataParams{
				FileID: fileID,
				Key:    md.Key,
				Value:  md.Value,
			}
			newMD, err := s.fileStore.CreateOrUpdateFileMetadata(ctx, argMD)
			if err != nil {
				return nil, logError(status.Errorf(codes.Internal, "failed to create or update file metadata: Err: %s", err))
			}
			newMDList = append(newMDList, newMD)
		}
	}
	return newMDList, nil
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
