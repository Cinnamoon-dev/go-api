package handlers

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"github.com/Cinnamoon-dev/go-api/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type DepartamentoHandler struct {
	service services.DepartamentoService
}

func NewDepartamentoHandler(s services.DepartamentoService) *DepartamentoHandler {
	return &DepartamentoHandler{service: s}
}

func (h *DepartamentoHandler) Insert(c echo.Context) error {
	var departamento models.Departamento

	if err := c.Bind(&departamento); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input."})
	}

	err := h.service.InsertDepartamento(departamento)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to insert trabalhador", "error": err.Error()})
	}

	return c.JSON(http.StatusCreated, departamento)
}

func (h *DepartamentoHandler) GetAll(c echo.Context) error {
	departamentos, err := h.service.GetAllDepartamentos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to fetch departamentos", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, departamentos)
}

func (h *DepartamentoHandler) GetByID(c echo.Context) error {
	idStr := c.Param("ID")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID must be a positive integer", "error": err.Error()})

	}
	departamento, err := h.service.GetDepartamentoByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Departamento not found", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, departamento)
}

func (h *DepartamentoHandler) Update(c echo.Context) error {
	var departamento models.Departamento
	if err := c.Bind(&departamento); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input", "error": err.Error()})
	}
	idStr := c.Param("ID")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID must be a positive integer", "error": err.Error()})
	}
	departamento.ID = id
	updatedDepartamento, err := h.service.UpdateDepartamento(departamento)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update departamento", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedDepartamento)
}

func (h *DepartamentoHandler) Delete(c echo.Context) error {
	idStr := c.Param("ID")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "ID must be a positive integer", "error": err.Error()})
	}
	departamento, err := h.service.DeleteDepartamento(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete departamento", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, departamento)
}
