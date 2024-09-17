package handlers

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"github.com/Cinnamoon-dev/go-api/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TrabalhadorHandler struct {
	service services.TrabalhadorService
}

func (h *TrabalhadorHandler) Insert(c echo.Context) error {
	var trabalhador models.Trabalhador
	if err := c.Bind(&trabalhador); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	err := h.service.InsertTrabalhador(trabalhador)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to insert trabalhador"})
	}

	return c.JSON(http.StatusCreated, trabalhador)
}

func NewTrabalhadorHandler(s services.TrabalhadorService) *TrabalhadorHandler {
	return &TrabalhadorHandler{service: s}
}

func (h *TrabalhadorHandler) GetAll(c echo.Context) error {
	trabalhadores, err := h.service.GetAllTrabalhadores()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch trabalhadores"})
	}
	return c.JSON(http.StatusOK, trabalhadores)
}
