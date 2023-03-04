package v1Name

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/islamic-name-generator-be/controller"
	"github.com/ranggarifqi/islamic-name-generator-be/src/my_error"
	"github.com/ranggarifqi/islamic-name-generator-be/src/name"
)

type nameHandler struct {
	nameService name.INameService
}

func SetupNameHandler(g *echo.Group, nameService name.INameService) {
	h := &nameHandler{
		nameService: nameService,
	}

	g.POST("/generate", h.Generate)
}

func (h *nameHandler) Generate(c echo.Context) error {
	payload := new(GenerateNameDTO)
	err := c.Bind(payload)
	if err != nil {
		apiError := controller.ConstructApiError(err)
		return c.JSON(apiError.Code, apiError)
	}

	err = c.Validate(payload)
	if err != nil {
		apiError := controller.ConstructApiError(my_error.NewBadRequestError(err.Error()))
		return c.JSON(apiError.Code, apiError)
	}

	result, err := h.nameService.GenerateName(name.GenerateNameDTO{
		Gender:              payload.Gender,
		ShouldUseMiddleName: *payload.ShouldUseMiddleName,
		ShouldUseLastName:   *payload.ShouldUseLastName,
	})
	if err != nil {
		return controller.ConstructApiError(err)
	}

	return c.JSON(http.StatusOK, result)
}
