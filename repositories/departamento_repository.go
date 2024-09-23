package repositories

import (
	"fmt"
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

	err := r.db.First(&existingDepartamento, departamento.ID).Error
	if err != nil {
		return departamento, err
	}

	existingDepartamento.Nome = departamento.Nome
	if err := r.db.Save(&existingDepartamento).Error; err != nil {
		return existingDepartamento, err
	}
	return existingDepartamento, nil
}

func (r *departamentoRepository) Delete(id uint64) (models.Departamento, error) {
	var departamento models.Departamento
	var trabalhador models.Trabalhador

	err := r.db.Where("id = ?", id).First(&departamento).Error
	if err != nil {
		return departamento, err
	}

	if err = r.db.Where("empresa_id = ?", departamento.ID).First(&trabalhador).Error; trabalhador.Nome != "" {
		return departamento, fmt.Errorf("Não é possível excluir o departamento com ID %d, pois há trabalhadores associados a ele", trabalhador.EmpresaID)
	}
	err = r.db.Delete(&departamento).Error
	if err != nil {
		return departamento, err
	}
	return departamento, nil
}
