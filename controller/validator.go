package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/ranggarifqi/islamic-name-generator-be/src/my_error"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return ConstructApiError(my_error.NewBadRequestError(err.Error()))
	}
	return nil
}
