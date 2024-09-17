package db

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Trabalhador{}, &models.Empresa{}, &models.Departamento{})
	return db, nil
}
