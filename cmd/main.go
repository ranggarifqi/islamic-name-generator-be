package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	v1 "github.com/ranggarifqi/islamic-name-generator-be/controller/v1"
	"github.com/ranggarifqi/islamic-name-generator-be/helper"
	"github.com/ranggarifqi/islamic-name-generator-be/mongodb"
	"github.com/ranggarifqi/islamic-name-generator-be/src/name"
)

func main() {
	helper.InitializeEnv("./.env")

	mongoDBClient, ctx, err := mongodb.Connect()
	if err != nil {
		panic(err)
	}
	defer mongodb.Disconnect(mongoDBClient, ctx)
	mongoDB := mongoDBClient.Database(os.Getenv("DB_NAME"))

	/* Setup Dependencies */
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	nameRepository := name.NewMongoRepository(ctx, mongoDB)
	nameService := name.NewService(nameRepository, randomizer)

	e := echo.New()

	// TODO: Setup CORS by reading from .env

	/* Setup Routing */
	v1Group := e.Group("/v1")
	v1.SetupHandler(v1Group, v1.V1Dependencies{
		NameService: nameService,
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
