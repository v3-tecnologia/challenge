package ierr

import "fmt"

type ValidationError struct {
	Err error
}

func (e *ValidationError) Error() string {
	return e.Err.Error()
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}

func NewValidationError(format string, args ...interface{}) error {
	return &ValidationError{
		Err: fmt.Errorf(format, args...),
	}
}
