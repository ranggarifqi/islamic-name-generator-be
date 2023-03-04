package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/islamic-name-generator-be/src/name"
)

type nameHandler struct {
	nameService name.INameService
}

func SetupNameHandler(g *echo.Group, dep V1Dependencies) {
	h := &nameHandler{
		nameService: dep.NameService,
	}

	g.POST("/", h.Generate)
}

func (h *nameHandler) Generate(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
