package server

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

type FileContentSaver interface {
	Save(filename string, filepath string, fileData bytes.Buffer) error
}

type DiskFileContentSaver struct {
	fileFolder string
}

func NewDiskFileContentSaver(fileFolder string) FileContentSaver {
	return &DiskFileContentSaver{fileFolder: fileFolder}
}

func (fs *DiskFileContentSaver) Save(filename string, filepath string, fileData bytes.Buffer) error {
	path := fmt.Sprintf("%s/%s", fs.fileFolder, filepath)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
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
