package models

type Trabalhador struct {
	ID           uint64         `gorm:"primaryKey"`
	Nome         string         `gorm:"not null;"`
	Cpf          string         `gorm:"not null;"`
	Empresa      []Empresa      `gorm:"ForeignKey:ID;"`
	Departamento []Departamento `gorm:"ForeignKey:ID;"`
}

type Empresa struct {
	ID          uint64 `gorm:"primaryKey"`
	RazaoSocial string `gorm:"not null;"`
	Cnpj        string `gorm:"not null;"`
}

type Departamento struct {
	ID   uint64 `gorm:"primaryKey"`
	Nome string `gorm:"not null;"`
}
