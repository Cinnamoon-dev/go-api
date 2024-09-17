package routes

import (
	"github.com/Cinnamoon-dev/go-api/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, trabalhadorHandler *handlers.TrabalhadorHandler) {
	e.GET("/trabalhadores", trabalhadorHandler.GetAll)
	e.POST("/trabalhadores", trabalhadorHandler.Insert)
}
