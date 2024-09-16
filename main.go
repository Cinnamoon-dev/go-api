package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type Trabalhador struct {
	ID           uint64         `gorm:"primaryKey"`
	Nome         string         `gorm:"not null;"`
	Cpf          string         `gorm:"not null;"`
	Empresa      []Empresa      `gorm:"ForeignKey:ID;"`
	Departamento []Departamento `gorm:"ForeignKey:ID;"`
}

type Empresa struct {
	ID          uint64 `gorm:"primaryKey"`
	RazaoSocial string `gorm:"not null;"`
	Cnpj        string `gorm:"not null;"`
}

type Departamento struct {
	ID   uint64 `gorm:"primaryKey"`
	Nome string `gorm:"not null;"`
}

func getDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func home(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"message": "hello"})
}

func main() {
	db, err := getDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Trabalhador{}, &Empresa{}, &Departamento{})

	e := echo.New()

	e.GET("/", home)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
