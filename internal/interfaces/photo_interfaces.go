package interfaces

import (
	"context"

	"github.com/KaiRibeiro/challenge/internal/models"
)

type Uploader interface {
	PutPhoto(ctx context.Context, filename string, image []byte) (string, error)
}

type FaceComparer interface {
	Compare(ctx context.Context, mac, filename string) (bool, error)
}

type PhotoService interface {
	AddPhoto(photo models.PhotoModel) (bool, error)
}
