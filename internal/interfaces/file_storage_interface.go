package interfaces

import "mime/multipart"

type FileStorageInterface interface {
	UploadFile(fileHeader *multipart.FileHeader, filename string) (string, error)
}
