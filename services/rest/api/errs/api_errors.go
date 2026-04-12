package errs

import "net/http"

type ApiError struct {
	err     error  `json:"-"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ApiError) Error() string {
	return e.err.Error()
}

func NewApiError(err error, status int, msg string) ApiError {
	return ApiError{
		err:     err,
		Status:  status,
		Message: msg,
	}
}

var InternalServerError = map[string]any{
	"status":  http.StatusInternalServerError,
	"message": "internal server error",
}
