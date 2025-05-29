package custom_errors

import "fmt"

type PhotoError struct {
	err    error
	status int
}

func NewPhotoError(err error, status int) *PhotoError {
	return &PhotoError{
		err:    err,
		status: status,
	}
}

func (e *PhotoError) Error() string {
	return fmt.Sprintf("Photo service failure: %s", e.err.Error())
}

func (e *PhotoError) Status() int {
	return e.status
}

func (e *PhotoError) Unwrap() error {
	return e.err
}

func IsPhotoError(err error) (*PhotoError, bool) {
	if photoErr, ok := err.(*PhotoError); ok {
		return photoErr, true
	}
	return nil, false
}
