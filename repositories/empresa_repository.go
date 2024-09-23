package repositories

import (
	"fmt"
	"github.com/Cinnamoon-dev/go-api/models"
	"gorm.io/gorm"
)

type EmpresaRepository interface {
	Insert(empresa models.Empresa) error
	GetAll() ([]models.Empresa, error)
	GetByCnpj(cnpj string) (models.Empresa, error)
	Update(empresa models.Empresa, cnpj string) (models.Empresa, error)
	Delete(cnpj string) (models.Empresa, error)
}

type empresaRepository struct {
	db *gorm.DB
}

func NewEmpresaRepository(db *gorm.DB) EmpresaRepository {
	return &empresaRepository{db: db}
}

func (r *empresaRepository) Insert(empresa models.Empresa) error {
	return r.db.Create(&empresa).Error
}

func (r *empresaRepository) GetAll() ([]models.Empresa, error) {
	var empresas []models.Empresa

	err := r.db.Find(&empresas).Error
	if err != nil {
		return nil, err
	}
	return empresas, nil
}

func (r *empresaRepository) GetByCnpj(cnpj string) (models.Empresa, error) {
	var empresa models.Empresa
	err := r.db.Where("cnpj = ?", cnpj).First(&empresa).Error
	if err != nil {
		return empresa, err
	}
	return empresa, nil
}

func (r *empresaRepository) Update(empresa models.Empresa, cnpjAntigo string) (models.Empresa, error) {
	var existingEmpresa models.Empresa

	if err := r.db.Where("cnpj = ?", cnpjAntigo).First(&existingEmpresa).Error; err != nil {
		return existingEmpresa, err
	}

	existingEmpresa.RazaoSocial = empresa.RazaoSocial
	existingEmpresa.Cnpj = empresa.Cnpj

	if err := r.db.Save(&existingEmpresa).Error; err != nil {
		return existingEmpresa, err
	}

	return existingEmpresa, nil
}

func (r *empresaRepository) Delete(cnpj string) (models.Empresa, error) {
	var trabalhador models.Trabalhador
	var empresa models.Empresa

	err := r.db.Where("cnpj = ?", cnpj).First(&empresa).Error
	if err != nil {
		return empresa, err
	}

	if err = r.db.Where("empresa_id = ?", empresa.ID).First(&trabalhador).Error; trabalhador.Nome != "" {
		return empresa, fmt.Errorf("Não é possível excluir a empresa com ID %d, pois há trabalhadores associados a ela", trabalhador.EmpresaID)
	}

	err = r.db.Delete(&empresa).Error
	if err != nil {
		return empresa, err
	}
	return empresa, nil
}
