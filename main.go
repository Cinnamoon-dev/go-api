package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Album struct {
	ID    uint64 `gorm:"primaryKey"`
	Title string
}

type Foto struct {
	ID      uint64 `gorm:"primaryKey"`
	AlbumID uint64 `gorm:"foreignKey:cu"`
	Uri     string

	Album Album
}

type Usuario struct {
	ID        uint64 `gorm:"primaryKey"`
	FotoID    uint64 `gorm:"primaryKey"`
	Nome      string
	Email     string
	Data_nasc time.Time
	Website   string
	Gender    string
	Telephone string

	Foto Foto
	//Seguidores []Seguidores `gorm:"foreignKey:UserID"`
	//Seguidos   []Seguidores `gorm:"foreignKey:FollowerID"`
}

//type Seguidores struct {
//	ID         uint64 `gorm:"primaryKey"`
//	SeguidoID  uint64 `gorm:"foreignKey:UsuarioID"`
//	SeguidorID uint64 `gorm:"foreignKey:SeguidorID"`
//
//	Usuario  Usuario `gorm:"foreignKey:UserID"`
//	Seguidor Usuario `gorm:"foreignKey:FollowerID"`
//}

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

	db.AutoMigrate(&Album{}, &Foto{}, &Usuario{})

	e := echo.New()

	e.GET("/", home)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
