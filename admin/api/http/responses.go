package http

type Response interface {
	Code() int
	Headers() map[string]string
	Empty() bool
}

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code"`
}
