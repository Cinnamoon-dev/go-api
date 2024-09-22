package handlers

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"github.com/Cinnamoon-dev/go-api/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EmpresaHandler struct {
	service services.EmpresaService
}

func NewEmpresaHandler(s services.EmpresaService) *EmpresaHandler {
	return &EmpresaHandler{service: s}
}

func (h *EmpresaHandler) Insert(c echo.Context) error {
	var empresa models.Empresa

	if err := c.Bind(&empresa); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input", "error": err.Error()})
	}

	err := h.service.InsertEmpresa(empresa)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to inser empresa", "error": err.Error()})
	}
	return c.JSON(http.StatusCreated, empresa)
}

func (h *EmpresaHandler) GetAll(c echo.Context) error {
	empresas, err := h.service.GetAllEmpresas()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to get empresas", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, empresas)
}

func (h *EmpresaHandler) GetByCnpj(c echo.Context) error {
	cnpj := c.Param("cnpj")
	empresa, err := h.service.GetEmpresaByCnpj(cnpj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to get empresa", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, empresa)
}

func (h *EmpresaHandler) Update(c echo.Context) error {
	var empresa models.Empresa
	if err := c.Bind(&empresa); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input", "error": err.Error()})
	}

	cnpj := c.Param("cnpj")

	updatedEmpresa, err := h.service.UpdateEmpresa(empresa, cnpj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update empresa", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedEmpresa)
}

func (h *EmpresaHandler) Delete(c echo.Context) error {
	var empresa models.Empresa

	cnpj := c.Param("cnpj")

	empresa, err := h.service.DeleteEmpresa(cnpj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete empresa", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, empresa)
}
