package controller

import (
	"errors"
	"net/http"
	"testing"

	"github.com/ranggarifqi/islamic-name-generator-be/src/my_error"
	"github.com/stretchr/testify/assert"
)

func Test_ConstructApiError(t *testing.T) {
	t.Run("Should infer my_error correctly", func(t *testing.T) {
		badRequest := my_error.NewBadRequestError("Bad Request!!!")

		err := ConstructApiError(badRequest)

		assert.Equal(t, http.StatusBadRequest, err.Code)
		assert.Equal(t, "Error 400: Bad Request!!!", err.Message)
	})

	t.Run("Should treat vanilla error as internal server error", func(t *testing.T) {
		randomErr := errors.New("Random Error")

		err := ConstructApiError(randomErr)

		assert.Equal(t, http.StatusInternalServerError, err.Code)
		assert.Equal(t, "Error 500: Random Error", err.Message)
	})
}
