package errorHelper

import (
	"fmt"
	"net/http"
)

type internalServerError struct {
	statusCode int
	errMessage string
}

func NewInternalServerError(errMessage string) error {
	return &internalServerError{
		statusCode: http.StatusInternalServerError,
		errMessage: errMessage,
	}
}

func (e *internalServerError) Error() string {
	return fmt.Sprintf("Error %v: %v", e.statusCode, e.errMessage)
}
