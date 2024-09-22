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

func NewTrabalhadorHandler(s services.TrabalhadorService) *TrabalhadorHandler {
	return &TrabalhadorHandler{service: s}
}

func (h *TrabalhadorHandler) Insert(c echo.Context) error {
	var trabalhador models.Trabalhador
	if err := c.Bind(&trabalhador); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input", "error": err.Error()})
	}

	trabalhador, err := h.service.InsertTrabalhador(trabalhador)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to insert trabalhador", "error": err.Error()})
	}

	return c.JSON(http.StatusCreated, trabalhador)
}

func (h *TrabalhadorHandler) GetAll(c echo.Context) error {
	trabalhadores, err := h.service.GetAllTrabalhadores()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to fetch trabalhadores", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, trabalhadores)
}

func (h *TrabalhadorHandler) GetByCpf(c echo.Context) error {
	cpf := c.Param("cpf")
	trabalhador, err := h.service.GetTrabalhadorByCpf(cpf)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Trabalhador not found", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, trabalhador)
}

func (h *TrabalhadorHandler) Update(c echo.Context) error {
	var trabalhador models.Trabalhador
	if err := c.Bind(&trabalhador); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input", "error": err.Error()})
	}

	cpf := c.Param("cpf")

	updatedTrabalhador, err := h.service.UpdateTrabalhador(trabalhador, cpf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update trabalhador", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedTrabalhador)
}

func (h *TrabalhadorHandler) Delete(c echo.Context) error {
	trabalhador, err := h.service.DeleteTrabalhador(c.Param("cpf"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete trabalhador", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, trabalhador)
}
