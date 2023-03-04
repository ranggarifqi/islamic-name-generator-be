package errorHelper

import (
	"fmt"
	"net/http"
)

type badRequestError struct {
	statusCode int
	errMessage string
}

func NewBadRequestError(errMessage string) error {
	return &badRequestError{
		statusCode: http.StatusBadRequest,
		errMessage: errMessage,
	}
}

func (e *badRequestError) Error() string {
	return fmt.Sprintf("Error %v: %v", e.statusCode, e.errMessage)
}
