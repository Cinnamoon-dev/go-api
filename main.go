package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Computers struct {
	ID    uint `gorm:"primaryKey"`
	Model string
	Brand string
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

	db.AutoMigrate(&Computers{})

	e := echo.New()

	e.GET("/", home)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
