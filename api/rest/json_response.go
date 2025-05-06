package rest

type JSONResponse struct {
	Error      bool        `json:"error"`
	Name       string      `json:"name,omitempty"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Action     string      `json:"action,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}
