package services

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"github.com/Cinnamoon-dev/go-api/repositories"
)

type DepartamentoService interface {
	InsertDepartamento(departamento models.Departamento) error
	GetAllDepartamentos() ([]models.Departamento, error)
	GetDepartamentoByID(id uint64) (models.Departamento, error)
	UpdateDepartamento(departamento models.Departamento) (models.Departamento, error)
	DeleteDepartamento(id uint64) (models.Departamento, error)
}

type departamentoService struct {
	repo repositories.DepartamentoRepository
}

func NewDepartamentoService(repo repositories.DepartamentoRepository) DepartamentoService {
	return &departamentoService{repo: repo}
}

func (s *departamentoService) InsertDepartamento(departamento models.Departamento) error {
	return s.repo.Insert(departamento)
}

func (s *departamentoService) GetAllDepartamentos() ([]models.Departamento, error) {
	return s.repo.GetAll()
}

func (s *departamentoService) GetDepartamentoByID(id uint64) (models.Departamento, error) {
	return s.repo.GetByID(id)
}

func (s *departamentoService) UpdateDepartamento(departamento models.Departamento) (models.Departamento, error) {
	return s.repo.Update(departamento)
}

func (s *departamentoService) DeleteDepartamento(id uint64) (models.Departamento, error) {
	return s.repo.Delete(id)
}
