package repositories

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"gorm.io/gorm"
)

type DepartamentoRepository interface {
	Insert(departamento models.Departamento) error
	GetAll() ([]models.Departamento, error)
	GetByID(id uint64) (models.Departamento, error)
	Update(departamento models.Departamento) (models.Departamento, error)
	Delete(id uint64) (models.Departamento, error)
}

type departamentoRepository struct {
	db *gorm.DB
}

func NewDepartamentoRepository(db *gorm.DB) DepartamentoRepository {
	return &departamentoRepository{db: db}
}

func (r *departamentoRepository) Insert(departamento models.Departamento) error {
	return r.db.Create(&departamento).Error
}

func (r *departamentoRepository) GetAll() ([]models.Departamento, error) {
	var departamentos []models.Departamento

	err := r.db.Find(&departamentos).Error
	if err != nil {
		return nil, err
	}

	return departamentos, nil

}

func (r *departamentoRepository) GetByID(id uint64) (models.Departamento, error) {
	var departamento models.Departamento

	err := r.db.First(&departamento, id).Error
	if err != nil {
		return departamento, err
	}

	return departamento, err
}

func (r *departamentoRepository) Update(departamento models.Departamento) (models.Departamento, error) {
	var existingDepartamento models.Departamento

	err := r.db.First(&departamento, departamento.ID).Error
	if err != nil {
		return departamento, err
	}
	err = r.db.Model(&existingDepartamento).Updates(map[string]interface{}{
		"nome": departamento.Nome,
	}).Error
	if err != nil {
		return departamento, err
	}
	return existingDepartamento, nil
}

func (r *departamentoRepository) Delete(id uint64) (models.Departamento, error) {
	var departamento models.Departamento

	err := r.db.Where("id = ?", id).First(&departamento).Error
	if err != nil {
		return departamento, err
	}
	err = r.db.Delete(&departamento).Error
	if err != nil {
		return departamento, err
	}
	return departamento, nil
}
