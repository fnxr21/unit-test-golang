package handlers

import (
	"net/http"
	"net/http/httptest"
	"unit-test/api/models"
	"testing"
	// "time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"errors"
)

type MockPenghuniRepository struct {
	Mock mock.Mock
}

var penghuniRepository = &MockPenghuniRepository{Mock: mock.Mock{}}
var penghuniHandler = HandlerGlobal{PenghuniRepository: penghuniRepository}

func (repository *MockPenghuniRepository) ClientGetNumber(handphone string) (models.Penghuni, error) {
	args := repository.Mock.Called(handphone)
	return args.Get(0).(models.Penghuni), args.Error(1)
}

func TestManagerCheckAuth(t *testing.T) {
	penghuni := models.Penghuni{
		ID:           3,
		PenghuniName: "richard",
		Handphone:    "6285925025265",
		Kategori:     1,
	}

	claims := jwt.MapClaims{}
	claims["id"] = 3
	claims["name"] = penghuni.PenghuniName
	claims["handphone"] = penghuni.Handphone
	claims["level"] = penghuni.Kategori

	t.Run("should return status 200 with still-active-manager", func(t *testing.T) {
		penghuniRepository.Mock.On("ClientGetNumber", penghuni.Handphone).Return(penghuni, nil).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("managerLogin", claims)

		err := penghuniHandler.ManagerCheckAuth(c)

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, `{"status":200,"data":"still-active-manager"}`, rec.Body.String())
		}
	})

	// if the user did not exist in db (ClientGetNumber returns empty penghuni)
	t.Run("should return status 401 with unauthorized", func(t *testing.T) {
		penghuniRepository.Mock.On("ClientGetNumber", penghuni.Handphone).Return(models.Penghuni{}, errors.New("some error message")).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("managerLogin", claims)

		err := penghuniHandler.ManagerCheckAuth(c)

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
			assert.JSONEq(t, `{"status":401,"message":"user-not-found"}`, rec.Body.String())
		}
	})
}
