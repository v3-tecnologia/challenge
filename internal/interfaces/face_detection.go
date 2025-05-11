package interfaces

import "context"

type FaceMatch struct {
	Similarity float64 `json:"similarity"`
	FaceID     string  `json:"faceId"`
}

type FaceDetector interface {
	HandleFaceRecognition(ctx context.Context, imageKey string) (*[]FaceMatch, error)
}
