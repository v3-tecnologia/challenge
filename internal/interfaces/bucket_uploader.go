package interfaces

import (
	"context"
	"io"
)

type BucketUploader interface {
	UploadAsync(ctx context.Context, file io.Reader, key string, ch chan<- string, errCh chan<- error)
}
