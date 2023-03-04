package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	v1 "github.com/ranggarifqi/islamic-name-generator-be/controller/v1"
	"github.com/ranggarifqi/islamic-name-generator-be/helper"
	"github.com/ranggarifqi/islamic-name-generator-be/mongodb"
)

func main() {
	helper.InitializeEnv("./.env")

	_, _, err := mongodb.Connect()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// TODO: Setup CORS by reading from .env

	v1Group := e.Group("/v1")
	v1.SetupHandler(v1Group, v1.V1Dependencies{})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
