package interfaces

import "context"

type FaceMatch struct {
	Similarity float64 `json:"similarity"`
	FaceID     string  `json:"faceId"`
}

type FaceDetector interface {
	HandleFaceRecognition(ctx context.Context, imageKey string) (*[]FaceMatch, error)

	//NOTE I'm leaving this method for demo purposes, it should not be used in production
	CreateCollection() (string, error)
}
