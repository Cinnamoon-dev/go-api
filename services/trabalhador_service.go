package services

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"github.com/Cinnamoon-dev/go-api/repositories"
)

type TrabalhadorService interface {
	InsertTrabalhador(trabalhador models.Trabalhador) error
	GetAllTrabalhadores() ([]models.Trabalhador, error)
	GetTrabalhadorByCpf(cpf string) (models.Trabalhador, error)
	UpdateTrabalhador(trabalhador models.Trabalhador) (models.Trabalhador, error)
	DeleteTrabalhador(cpf string) (models.Trabalhador, error)
}

type trabalhadorService struct {
	repo repositories.TrabalhadorRepository
}

func NewTrabalhadorService(repo repositories.TrabalhadorRepository) TrabalhadorService {
	return &trabalhadorService{repo: repo}
}

func (s *trabalhadorService) InsertTrabalhador(trabalhador models.Trabalhador) error {
	return s.repo.Insert(trabalhador)
}

func (s *trabalhadorService) GetAllTrabalhadores() ([]models.Trabalhador, error) {
	return s.repo.GetAll()
}

func (s *trabalhadorService) GetTrabalhadorByCpf(cpf string) (models.Trabalhador, error) {
	return s.repo.GetByCpf(cpf)
}

func (s *trabalhadorService) UpdateTrabalhador(trabalhador models.Trabalhador) (models.Trabalhador, error) {
	return s.repo.Update(trabalhador)
}

func (s *trabalhadorService) DeleteTrabalhador(cpf string) (models.Trabalhador, error) {
	return s.repo.Delete(cpf)
}
