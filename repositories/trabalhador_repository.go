package repositories

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"gorm.io/gorm"
)

type TrabalhadorRepository interface {
	GetAll() ([]models.Trabalhador, error)
	Insert(trabalhador models.Trabalhador) error
}

type trabalhadorRepository struct {
	db *gorm.DB
}

func (r *trabalhadorRepository) Insert(trabalhador models.Trabalhador) error {
	result := r.db.Create(&trabalhador)
	return result.Error
}

func NewTrabalhadorRepository(db *gorm.DB) TrabalhadorRepository {
	return &trabalhadorRepository{db: db}
}

func (r *trabalhadorRepository) GetAll() ([]models.Trabalhador, error) {
	var trabalhadores []models.Trabalhador
	err := r.db.Preload("Empresa").Preload("Departamento").Find(&trabalhadores).Error
	return trabalhadores, err
}
