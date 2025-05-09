package interfaces

import "io"

type BucketUploader interface {
	Upload(file io.Reader, key string, ch chan<- string)
}
