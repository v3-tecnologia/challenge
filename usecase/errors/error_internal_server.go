package appError

type ErrorInternalServer struct {
	Code    int
	Message string
	Action  string
}

func NewErrorInternalServer(message string, action string) *ErrorInternalServer {
	return &ErrorInternalServer{
		Code:    500,
		Message: message,
		Action:  action,
	}
}

func (e *ErrorInternalServer) Error() string     { return e.Message }
func (e *ErrorInternalServer) GetAction() string { return e.Action }
func (e *ErrorInternalServer) GetCode() int      { return e.Code }
