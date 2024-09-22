package main

import (
	"github.com/Cinnamoon-dev/go-api/db"
	"github.com/Cinnamoon-dev/go-api/handlers"
	"github.com/Cinnamoon-dev/go-api/repositories"
	"github.com/Cinnamoon-dev/go-api/routes"
	"github.com/Cinnamoon-dev/go-api/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

// inicializacao da aplicacao
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.Init()
	if err != nil {
		panic("Failed to connect to database")
	}

	trabalhadorRepo := repositories.NewTrabalhadorRepository(database)
	trabalhadorService := services.NewTrabalhadorService(trabalhadorRepo)
	trabalhadorHandler := handlers.NewTrabalhadorHandler(trabalhadorService)

	empresaRepo := repositories.NewEmpresaRepository(database)
	empresaService := services.NewEmpresaService(empresaRepo)
	empresaHandler := handlers.NewEmpresaHandler(empresaService)

	departamentoRepo := repositories.NewDepartamentoRepository(database)
	departamentoService := services.NewDepartamentoService(departamentoRepo)
	departamentoHandler := handlers.NewDepartamentoHandler(departamentoService)

	port := os.Getenv("PORT")

	e := echo.New()

	routes.TrabalhadorRegisterRoutes(e, trabalhadorHandler)
	routes.EmpresaRegisterRoutes(e, empresaHandler)
	routes.DepartamentoRegisterRoutes(e, departamentoHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
