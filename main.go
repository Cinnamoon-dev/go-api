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

	port := os.Getenv("PORT")

	e := echo.New()
	routes.RegisterRoutes(e, trabalhadorHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
