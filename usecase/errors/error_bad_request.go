package appError

type ErrorBadRequest struct {
	Code    int
	Message string
	Action  string
}

func NewErrorBadRequest(message string, action string) *ErrorBadRequest {
	return &ErrorBadRequest{
		Code:    400,
		Message: message,
		Action:  action,
	}
}

func (e *ErrorBadRequest) Error() string     { return e.Message }
func (e *ErrorBadRequest) GetAction() string { return e.Action }
func (e *ErrorBadRequest) GetCode() int      { return e.Code }
