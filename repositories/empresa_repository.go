package repositories

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"gorm.io/gorm"
)

type EmpresaRepository interface {
	Insert(empresa models.Empresa) error
	GetAll() ([]models.Empresa, error)
	GetByCnpj(cnpj string) (models.Empresa, error)
	Update(empresa models.Empresa) (models.Empresa, error)
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

func (r *empresaRepository) Update(empresa models.Empresa) (models.Empresa, error) {
	var existingEmpresa models.Empresa

	err := r.db.Where("cnpj = ?", empresa.Cnpj).First(&existingEmpresa).Error
	if err != nil {
		return empresa, err
	}

	err = r.db.Model(&existingEmpresa).Updates(map[string]interface{}{
		"razao_social": empresa.RazaoSocial,
		"cnpj":         empresa.Cnpj,
	}).Error

	if err != nil {
		return models.Empresa{}, err
	}
	return empresa, nil
}

func (r *empresaRepository) Delete(cnpj string) (models.Empresa, error) {
	var empresa models.Empresa

	err := r.db.Where("cnpj = ?", cnpj).First(&empresa).Error
	if err != nil {
		return empresa, err
	}

	err = r.db.Delete(&empresa).Error
	if err != nil {
		return empresa, err
	}
	return empresa, nil
}
