package middleware

import (
	"log"
	"unit-test/api/handlers"
	jwtToken "unit-test/api/pkg/jwt"

	"net/http"
	"strings"

	"fmt"

	"github.com/labstack/echo/v4"
)

// Declare Result struct here ...
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Create Auth function here ...
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, handlers.ErrorResult{Status: http.StatusBadRequest, Message: "unauthorized-token"})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		fmt.Print("CLAIM : ", claims)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, handlers.ErrorResult{Status: http.StatusUnauthorized, Message: "unauthorized"})
		}

		level, ok := claims["level"].(float64)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}
		if level != 1 {

			status, _ := claims["status"].(string)
			if status == "manager" {
				log.Println(claims, "manager====")
				c.Set("managerLogin", claims)
				return next(c)
			}
		}

		return c.JSON(http.StatusUnauthorized, "unauthorized-not-found")
	}
}
