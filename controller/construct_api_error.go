package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/islamic-name-generator-be/src/my_error"
)

func ConstructApiError(err error) *echo.HTTPError {
	myError, ok := err.(my_error.Error)

	if ok {
		statusCode := myError.GetStatusCode()
		return echo.NewHTTPError(statusCode, myError.Error())
	}

	errMessage := fmt.Sprintf("Error 500: %v", err.Error())
	return echo.NewHTTPError(http.StatusInternalServerError, errMessage)
}
