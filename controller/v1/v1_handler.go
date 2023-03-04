package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/islamic-name-generator-be/src/name"
)

type V1Dependencies struct {
	NameService name.INameService
}

func SetupHandler(g *echo.Group, dep V1Dependencies) {
	nameRoute := g.Group("/name")

	SetupNameHandler(nameRoute, dep)
}
