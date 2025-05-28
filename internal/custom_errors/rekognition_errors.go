package custom_errors

import "fmt"

type RekognitionError struct {
	err    error
	status int
}

func NewRekognitionError(err error, status int) *RekognitionError {
	return &RekognitionError{
		err:    err,
		status: status,
	}
}

func (e *RekognitionError) Error() string {
	return fmt.Sprintf("AWS Rekognition error: %s", e.err.Error())
}

func (e *RekognitionError) Status() int {
	return e.status
}

func (e *RekognitionError) Unwrap() error {
	return e.err
}

func IsRekognitionError(err error) (*RekognitionError, bool) {
	if rekoErr, ok := err.(*RekognitionError); ok {
		return rekoErr, true
	}
	return nil, false
}
