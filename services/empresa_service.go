package services

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"github.com/Cinnamoon-dev/go-api/repositories"
)

type EmpresaService interface {
	InsertEmpresa(empresa models.Empresa) error
	GetAllEmpresas() ([]models.Empresa, error)
	GetEmpresaByCnpj(cnpj string) (models.Empresa, error)
	UpdateEmpresa(empresa models.Empresa) (models.Empresa, error)
	DeleteEmpresa(cnpj string) (models.Empresa, error)
}

type empresaService struct {
	repo repositories.EmpresaRepository
}

func NewEmpresaService(repo repositories.EmpresaRepository) EmpresaService {
	return &empresaService{repo: repo}
}

func (s *empresaService) InsertEmpresa(empresa models.Empresa) error {
	return s.repo.Insert(empresa)
}

func (s *empresaService) GetAllEmpresas() ([]models.Empresa, error) {
	return s.repo.GetAll()
}

func (s *empresaService) GetEmpresaByCnpj(cnpj string) (models.Empresa, error) {
	return s.repo.GetByCnpj(cnpj)
}

func (s *empresaService) UpdateEmpresa(empresa models.Empresa) (models.Empresa, error) {
	return s.repo.Update(empresa)
}

func (s *empresaService) DeleteEmpresa(cnpj string) (models.Empresa, error) {
	return s.repo.Delete(cnpj)
}
