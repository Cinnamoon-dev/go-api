package services

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"github.com/Cinnamoon-dev/go-api/repositories"
)

type TrabalhadorService interface {
	GetAllTrabalhadores() ([]models.Trabalhador, error)
	InsertTrabalhador(trabalhador models.Trabalhador) error
}

type trabalhadorService struct {
	repo repositories.TrabalhadorRepository
}

func (s *trabalhadorService) InsertTrabalhador(trabalhador models.Trabalhador) error {
	return s.repo.Insert(trabalhador)
}

func NewTrabalhadorService(repo repositories.TrabalhadorRepository) TrabalhadorService {
	return &trabalhadorService{repo: repo}
}

func (s *trabalhadorService) GetAllTrabalhadores() ([]models.Trabalhador, error) {
	return s.repo.GetAll()
}
