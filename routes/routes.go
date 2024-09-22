package routes

import (
	"github.com/Cinnamoon-dev/go-api/handlers"
	"github.com/labstack/echo/v4"
)

func TrabalhadorRegisterRoutes(e *echo.Echo, trabalhadorHandler *handlers.TrabalhadorHandler) {
	// Rota para criar um novo trabalhador
	e.POST("/trabalhadores", trabalhadorHandler.Insert)

	// Rota para listar todos os trabalhadores
	e.GET("/trabalhadores", trabalhadorHandler.GetAll)

	// Rota para listar um trabalhador através do seu CPF
	e.GET("/trabalhadores/:cpf", trabalhadorHandler.GetByCpf)

	// Rota para atualizar os dados de um trabalhador existente através do seu CPF
	e.PUT("/trabalhadores/:cpf", trabalhadorHandler.Update)

	// Rota para deletar um trabalhador através do CPF
	e.DELETE("/trabalhadores/:cpf", trabalhadorHandler.Delete)
}

func EmpresaRegisterRoutes(e *echo.Echo, empresaHandler *handlers.EmpresaHandler) {
	// Rota para criar uma nova empresa
	e.POST("/empresas", empresaHandler.Insert)

	// Rota para listar todas as empresas
	e.GET("/empresas", empresaHandler.GetAll)

	// Rota para listar uma empresa através do seu CNPJ
	e.GET("/empresas/:cnpj", empresaHandler.GetByCnpj)

	// Rota para atualizar os dados de uma empresa existente através do seu CNPJ
	e.PUT("/empresas/:cnpj", empresaHandler.Update)

	// Rota para deletar uma empresa através do CNPJ
	e.DELETE("/empresas/:cnpj", empresaHandler.Delete)
}

func DepartamentoRegisterRoutes(e *echo.Echo, departamentoHandler *handlers.DepartamentoHandler) {
	// Rota para criar um novo departamento
	e.POST("/departamentos", departamentoHandler.Insert)

	// Rota para listar todos os trabalhadores
	e.GET("/departamentos", departamentoHandler.GetAll)

	// Rota para listar um departamento através do seu ID
	e.GET("/departamentos/:ID", departamentoHandler.GetByID)

	// Rota para atualizar os dados de um departamento existente através do seu ID
	e.PUT("/departamentos/:ID", departamentoHandler.Update)

	// Rota para deletar um departamento através do ID
	e.DELETE("/departamentos/:ID", departamentoHandler.Delete)
}
