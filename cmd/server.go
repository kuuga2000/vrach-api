package main

import (
	"net/http"
	customerHandler "vrachapi/cmd/customer/handler"
	"vrachapi/configs"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.POST(configs.Version()+"/register", customerHandler.Register)

	configs.InitDB()
	e.Logger.Fatal(e.Start(":80"))
}
