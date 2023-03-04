package v1Name

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/islamic-name-generator-be/controller"
	"github.com/ranggarifqi/islamic-name-generator-be/src/name"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Generate(t *testing.T) {
	// Setup
	randomizer := rand.New(rand.NewSource(1))
	nameRepository := name.NewMockINameRepository(t)
	nameService := name.NewService(nameRepository, randomizer)

	e := echo.New()
	e.Validator = &controller.CustomValidator{Validator: validator.New()}

	payloadStr := `{"gender": "IKHWAN", "shouldUseMiddleName": false, "shouldUseLastName": true}`

	req := httptest.NewRequest(http.MethodPost, "/v1/name/generate", strings.NewReader(payloadStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	h := &nameHandler{
		nameService: nameService,
	}

	t.Run("Should generate name successfully", func(t *testing.T) {
		findByResults := []name.Name{
			{
				ID:        "firstName1",
				Name:      "andi",
				Gender:    name.IKHWAN,
				NameTypes: []name.NameType{name.FIRST_NAME},
			},
			{
				ID:        "middleName1",
				Name:      "putra",
				Gender:    name.IKHWAN,
				NameTypes: []name.NameType{name.MIDDLE_NAME},
			},
			{
				ID:        "lastName1",
				Name:      "sulistyo",
				Gender:    name.IKHWAN,
				NameTypes: []name.NameType{name.LAST_NAME},
			},
		}

		nameRepository.On("FindBy", mock.Anything).Return(&findByResults, nil)

		err := h.Generate(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		expectedJSON := `{"FIRST_NAME":{"ID":"firstName1","Name":"andi","Gender":"IKHWAN","NameTypes":["FIRST_NAME"],"Meanings":null,"CreatedAt":"0001-01-01T00:00:00Z"},"LAST_NAME":{"ID":"lastName1","Name":"sulistyo","Gender":"IKHWAN","NameTypes":["LAST_NAME"],"Meanings":null,"CreatedAt":"0001-01-01T00:00:00Z"}}
`

		assert.Equal(t, expectedJSON, rec.Body.String())
	})
}
