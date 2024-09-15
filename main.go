package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK, "home")
}

func main() {
	e := echo.New()

	e.GET("/", home)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
