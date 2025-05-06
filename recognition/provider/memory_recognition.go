package recognition

import (
	"context"

	"github.com/mkafonso/go-cloud-challenge/recognition"
)

type InMemoryFaceRecognition struct{}

func NewInMemoryFaceRecognition() recognition.FaceRecognitionService {
	return &InMemoryFaceRecognition{}
}

func (r *InMemoryFaceRecognition) CompareWithHistory(ctx context.Context, photoPath, deviceID string) (bool, error) {
	return true, nil
}
