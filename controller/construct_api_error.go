package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/islamic-name-generator-be/src/my_error"
)

func ConstructApiError(c echo.Context, err error) error {
	myError, ok := err.(my_error.Error)

	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return echo.NewHTTPError(myError.GetStatusCode(), myError.Error())
}
