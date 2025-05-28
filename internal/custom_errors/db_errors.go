package custom_errors

import "fmt"

type DBError struct {
	err    error
	status int
}

func NewDBError(err error, status int) *DBError {
	return &DBError{
		err:    err,
		status: status,
	}
}

func (e *DBError) Error() string {
	return fmt.Sprintf("Database error: %s", e.err.Error())
}

func (e *DBError) Status() int {
	return e.status
}

func (e *DBError) Unwrap() error {
	return e.err
}

func IsDBError(err error) (*DBError, bool) {
	if dbErr, ok := err.(*DBError); ok {
		return dbErr, true
	}
	return nil, false
}
