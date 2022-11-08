package response

import (
	"net/http"
	"time"
)

type Response struct {
	Timestamp time.Time `json:"timestamp"`
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
}

func NewResponse(code int, message string) (int, Response) {
	return code, Response{
		Timestamp: time.Now(),
		Code:      code,
		Status:    http.StatusText(code),
		Message:   message,
	}
}

func NewOkResponse(message string) Response {
	return Response{
		Timestamp: time.Now(),
		Code:      http.StatusOK,
		Status:    http.StatusText(http.StatusOK),
		Message:   message,
	}
}

