package repositories

import (
	"github.com/Cinnamoon-dev/go-api/models"
	"gorm.io/gorm"
)

// Registra os metodos que iram compor o trabalhador Repositorio obrigando a serem implementados
type TrabalhadorRepository interface {
	Insert(trabalhador models.Trabalhador) error
	GetAll() ([]models.Trabalhador, error)
	GetByCpf(cpf string) (models.Trabalhador, error)
	Update(trabalhador models.Trabalhador) (models.Trabalhador, error)
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
func (r *trabalhadorRepository) Insert(trabalhador models.Trabalhador) error {
	return r.db.Create(&trabalhador).Error
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

func (r *trabalhadorRepository) Update(trabalhador models.Trabalhador) (models.Trabalhador, error) {
	var existingTrabalhador models.Trabalhador

	// Buscar o trabalhador existente com base no ID e do identificador único)
	err := r.db.Where("cpf = ?", trabalhador.Cpf).First(&existingTrabalhador).Error
	if err != nil {
		return trabalhador, err // Retorna erro se o trabalhador não for encontrado
	}

	// Atualizar apenas os campos fornecidos
	err = r.db.Model(&existingTrabalhador).Updates(map[string]interface{}{
		"cpf":             trabalhador.Cpf,
		"nome":            trabalhador.Nome,
		"empresa_id":      trabalhador.EmpresaID,      // Atualiza apenas o ID da empresa, não o objeto completo
		"departamento_id": trabalhador.DepartamentoID, // Atualiza apenas o ID do departamento, não o objeto completo
	}).Error

	if err != nil {
		return models.Trabalhador{}, err
	}

	// Retorna o trabalhador atualizado
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
