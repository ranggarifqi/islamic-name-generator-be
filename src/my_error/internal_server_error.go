package my_error

import (
	"fmt"
	"net/http"
)

type internalServerError struct {
	statusCode int
	errMessage string
}

func NewInternalServerError(errMessage string) MyError {
	return &internalServerError{
		statusCode: http.StatusInternalServerError,
		errMessage: errMessage,
	}
}

func (e *internalServerError) Error() string {
	return fmt.Sprintf("Error %v: %v", e.statusCode, e.errMessage)
}

func (e *internalServerError) GetStatusCode() int {
	return e.statusCode
}
