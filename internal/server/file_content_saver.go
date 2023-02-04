package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

type FileContentSaver interface {
	Save(ctx context.Context, filename string, filepath string, fileData bytes.Buffer) error
	Delete(ctx context.Context, filename string, filepath string) error
	Rename(ctx context.Context, oldName string, newName string, pathname string) error
}

type DiskFileContentSaver struct {
	fileFolder string
}

func NewDiskFileContentSaver(fileFolder string) FileContentSaver {
	return &DiskFileContentSaver{fileFolder: fileFolder}
}

func (fs *DiskFileContentSaver) Save(ctx context.Context, filename string, filepath string, fileData bytes.Buffer) error {
	path := fmt.Sprintf("%s/%s", fs.fileFolder, filepath)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("cannot create directory: %w", err)
		}
	}

	filePath := fmt.Sprintf("%s/%s", path, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}

	_, err = fileData.WriteTo(file)
	if err != nil {
		return fmt.Errorf("cannot write image to file: %w", err)
	}
	return nil
}

func (fs *DiskFileContentSaver) Delete(ctx context.Context, filename string, filepath string) error {
	file := fmt.Sprintf("%s/%s/%s", fs.fileFolder, filepath, filename)
	log.Debug().Msg(file)

	err := os.Remove(file)
	if err != nil {
		return nil
	}

	return nil
}

func (fs *DiskFileContentSaver) Rename(
	ctx context.Context,
	oldName string,
	newName string,
	pathname string,
) error {
	originFile := fmt.Sprintf("%s/%s/%s", fs.fileFolder, pathname, oldName)
	newFile := fmt.Sprintf("%s/%s/%s", fs.fileFolder, pathname, newName)
	err := os.Rename(originFile, newFile)
	if err != nil {
		return err
	}

	return nil
}
