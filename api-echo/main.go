package main

import (
	"net/http"
	"unit-test/api-echo/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

		// Inisialisasi handler
		h := handler.NewHandler()

		// Routing
		e.POST("/users", h.CreateUser)
		e.GET("/users/:email", h.GetUser)

	e.Logger.Fatal(e.Start(":1323"))
}
