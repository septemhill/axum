package axum

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	httpStatusCode int
	errCode        int
	errMessage     string
}

func (er *ErrorResponse) Pack() (*Response, error) {
	return &Response{
		Body: map[string]interface{}{
			"error": er.errMessage,
			"code":  er.errCode,
		},
		HTTPStatusCode: er.httpStatusCode,
	}, nil
}

func InvalidParameter(errCode int, msg string, args ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		httpStatusCode: http.StatusBadRequest,
		errCode:        errCode,
		errMessage:     "invalid parameter: " + fmt.Sprintf(msg, args...),
	}
}

func PermissionDenied(errCode int, msg string, args ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		httpStatusCode: http.StatusForbidden,
		errCode:        errCode,
		errMessage:     "permission denied: " + fmt.Sprintf(msg, args...),
	}
}

func Unauthorized(errCode int, msg string, args ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		httpStatusCode: http.StatusUnauthorized,
		errCode:        errCode,
		errMessage:     "unauthorized: " + fmt.Sprintf(msg, args...),
	}
}
