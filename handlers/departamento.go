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
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input."})
	}

	err := h.service.InsertDepartamento(departamento)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to insert trabalhador"})
	}

	return c.JSON(http.StatusCreated, departamento)
}

func (h *DepartamentoHandler) GetAll(c echo.Context) error {
	departamentos, err := h.service.GetAllDepartamentos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch departamentos"})
	}
	return c.JSON(http.StatusOK, departamentos)
}

func (h *DepartamentoHandler) GetByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID must be a positive integer"})

	}
	departamento, err := h.service.GetDepartamentoByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Departamento not found"})
	}
	return c.JSON(http.StatusOK, departamento)
}

func (h *DepartamentoHandler) Update(c echo.Context) error {
	var departamento models.Departamento
	if err := c.Bind(&departamento); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input."})
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID must be a positive integer"})
	}
	departamento.ID = id
	updatedDepartamento, err := h.service.UpdateDepartamento(departamento)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update departamento"})
	}
	return c.JSON(http.StatusOK, updatedDepartamento)
}

func (h *DepartamentoHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID must be a positive integer"})
	}
	departamento, err := h.service.DeleteDepartamento(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete departamento"})
	}
	return c.JSON(http.StatusOK, departamento)
}
