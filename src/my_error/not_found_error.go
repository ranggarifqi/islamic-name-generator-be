package my_error

import (
	"fmt"
	"net/http"
)

type notFoundError struct {
	statusCode int
	errMessage string
}

func NewNotFoundError(errMessage string) Error {
	return &notFoundError{
		statusCode: http.StatusNotFound,
		errMessage: errMessage,
	}
}

func (e *notFoundError) Error() string {
	return fmt.Sprintf("Error %v: %v", e.statusCode, e.errMessage)
}

func (e *notFoundError) GetStatusCode() int {
	return e.statusCode
}
