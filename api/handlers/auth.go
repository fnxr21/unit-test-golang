package handlers

import (
	"log"

	"net/http"
	"unit-test/api/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type HandlerGlobal struct {
	PenghuniRepository repositories.Penghuni
}

func Handler(handler HandlerGlobal) *HandlerGlobal {
	return &HandlerGlobal{
		PenghuniRepository: handler.PenghuniRepository,
	}
}

func (h *HandlerGlobal) CheckAuth(c echo.Context) error {

	userClaims := c.Get("UserLogin").(jwt.MapClaims)
	handphone, _ := userClaims["handphone"].(string)

	data := "not-active"
	FInduser, err := h.PenghuniRepository.ClientGetNumber(handphone)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResult{Status: http.StatusUnauthorized, Message: "user-not-found"})

	}
	if FInduser.Handphone != "" {
		data = "still-active"
		log.Println("user")
	}

	return c.JSON(http.StatusOK, SuccessResult{Status: http.StatusOK, Data: data})

}

//dt

type SuccessResult struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
type ErrorResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
