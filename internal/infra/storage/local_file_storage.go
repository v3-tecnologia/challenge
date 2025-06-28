package local_file_storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalPhotoStorage struct {
	BasePath string
}

func NewLocalPhotoStorage(basePath string) *LocalPhotoStorage {
	return &LocalPhotoStorage{BasePath: basePath}
}

func (s *LocalPhotoStorage) UploadFile(fileHeader *multipart.FileHeader, filename string) (string, error) {
	err := os.MkdirAll(s.BasePath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	fullPath := filepath.Join(s.BasePath, filename)

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file on disk: %w", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}

	return fullPath, nil
}
