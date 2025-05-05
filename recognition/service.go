package recognition

import "context"

type FaceRecognitionService interface {
	CompareWithHistory(ctx context.Context, photoPath, deviceID string) (bool, error)
}
