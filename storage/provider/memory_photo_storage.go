package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type InMemoryPhotoStorage struct{}

func NewInMemoryPhotoStorage() *InMemoryPhotoStorage {
	return &InMemoryPhotoStorage{}
}

func (s *InMemoryPhotoStorage) Store(ctx context.Context, deviceID string, fileBytes []byte) (string, error) {
	outputDir := "./uploads"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return "", err
	}

	safeDeviceID := strings.ReplaceAll(deviceID, ":", "_")
	filePath := filepath.Join(outputDir, fmt.Sprintf("photo_%s.jpg", safeDeviceID))

	if err := os.WriteFile(filePath, fileBytes, 0644); err != nil {
		return "", err
	}

	return filePath, nil
}
