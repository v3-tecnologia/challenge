package custom_errors

import "fmt"

type S3Error struct {
	err    error
	status int
}

func NewS3Error(err error, status int) *S3Error {
	return &S3Error{
		err:    err,
		status: status,
	}
}

func (e *S3Error) Error() string {
	return fmt.Sprintf("AWS s3 error: %s", e.err.Error())
}

func (e *S3Error) Status() int {
	return e.status
}

func (e *S3Error) Unwrap() error {
	return e.err
}

func IsS3Error(err error) (*S3Error, bool) {
	if s3Err, ok := err.(*S3Error); ok {
		return s3Err, true
	}
	return nil, false
}
