package repositories

import (
	"fmt"
	"github.com/Cinnamoon-dev/go-api/models"
	"gorm.io/gorm"
)

// Registra os metodos que iram compor o trabalhador Repositorio obrigando a serem implementados
type TrabalhadorRepository interface {
	Insert(trabalhador models.Trabalhador) (models.Trabalhador, error)
	GetAll() ([]models.Trabalhador, error)
	GetByCpf(cpf string) (models.Trabalhador, error)
	Update(trabalhador models.Trabalhador, cpf string) (models.Trabalhador, error)
	Delete(cpf string) (models.Trabalhador, error)
}

// struct que ira receber a instancia do banco de dados
type trabalhadorRepository struct {
	db *gorm.DB
}

func NewTrabalhadorRepository(db *gorm.DB) TrabalhadorRepository {
	return &trabalhadorRepository{db: db}
}

// Metodo de insert do trabalhador, recebendo um trabalhador como parametro, chamando a conexao do banco relacionado ao repositorio do trabalhador e inserindo os valores no banco
func (r *trabalhadorRepository) Insert(trabalhador models.Trabalhador) (models.Trabalhador, error) {
	var existingEmpresa models.Empresa
	var existingDepartamento models.Departamento

	if trabalhador.EmpresaID != 0 {
		if err := r.db.Where("id = ?", trabalhador.EmpresaID).First(&existingEmpresa).Error; err != nil {
			return trabalhador, fmt.Errorf("empresa com ID %d não encontrada", trabalhador.EmpresaID)
		}
		trabalhador.EmpresaID = existingEmpresa.ID
	}

	if trabalhador.DepartamentoID != 0 {
		if err := r.db.Where("id = ?", trabalhador.DepartamentoID).First(&existingDepartamento).Error; err != nil {
			return trabalhador, fmt.Errorf("departamento com ID %d não encontrado", trabalhador.DepartamentoID)
		}
		trabalhador.DepartamentoID = existingDepartamento.ID
	}

	if err := r.db.Create(&trabalhador).Error; err != nil {
		return trabalhador, err
	}
	err := r.db.Preload("Departamento").Preload("Empresa").Where("cpf = ?", trabalhador.Cpf).First(&trabalhador).Error

	return trabalhador, err
}

func (r *trabalhadorRepository) GetAll() ([]models.Trabalhador, error) {
	var trabalhadores []models.Trabalhador

	// Carregar todos os trabalhadores com suas respectivas empresas e departamentos
	err := r.db.Preload("Empresa").Preload("Departamento").Find(&trabalhadores).Error
	if err != nil {
		return nil, err
	}

	return trabalhadores, nil
}

// Retorna o trabalhador de acordo com o CPF
func (r *trabalhadorRepository) GetByCpf(cpf string) (models.Trabalhador, error) {
	var trabalhador models.Trabalhador
	err := r.db.Preload("Departamento").Preload("Empresa").Where("cpf = ?", cpf).First(&trabalhador).Error
	if err != nil {
		return trabalhador, err
	}
	return trabalhador, err
}

func (r *trabalhadorRepository) Update(trabalhador models.Trabalhador, cpfAntigo string) (models.Trabalhador, error) {
	var existingTrabalhador models.Trabalhador
	var existingEmpresa models.Empresa
	var existingDepartamento models.Departamento

	err := r.db.Where("cpf = ?", cpfAntigo).First(&existingTrabalhador).Error
	if err != nil {
		return trabalhador, fmt.Errorf("trabalhador com CPF %s não encontrado", cpfAntigo)
	}

	if trabalhador.EmpresaID != 0 {
		if err := r.db.Where("id = ?", trabalhador.EmpresaID).First(&existingEmpresa).Error; err != nil {
			return trabalhador, fmt.Errorf("empresa com ID %d não encontrada", trabalhador.EmpresaID)
		}
		existingTrabalhador.EmpresaID = trabalhador.EmpresaID
	}

	if trabalhador.DepartamentoID != 0 {
		if err := r.db.Where("id = ?", trabalhador.DepartamentoID).First(&existingDepartamento).Error; err != nil {
			return trabalhador, fmt.Errorf("departamento com ID %d não encontrado", trabalhador.DepartamentoID)
		}
		existingTrabalhador.DepartamentoID = trabalhador.DepartamentoID
	}

	existingTrabalhador.Nome = trabalhador.Nome
	existingTrabalhador.Cpf = trabalhador.Cpf

	err = r.db.Save(&existingTrabalhador).Error
	if err != nil {
		return trabalhador, fmt.Errorf("falha ao atualizar trabalhador: %v", err)
	}
	err = r.db.Preload("Departamento").Preload("Empresa").Where("cpf = ?", existingTrabalhador.Cpf).First(&existingTrabalhador).Error

	return existingTrabalhador, nil
}

func (r *trabalhadorRepository) Delete(cpf string) (models.Trabalhador, error) {
	var trabalhador models.Trabalhador

	// Primeiro buscar o trabalhador com a empresa e o departamento associados
	err := r.db.Where("cpf = ?", cpf).Preload("Empresa").Preload("Departamento").First(&trabalhador).Error
	if err != nil {
		return models.Trabalhador{}, err // Retorna erro caso o trabalhador não seja encontrado
	}

	// Deletar o trabalhador
	err = r.db.Delete(&trabalhador).Error
	if err != nil {
		return models.Trabalhador{}, err
	}

	return trabalhador, nil
}
