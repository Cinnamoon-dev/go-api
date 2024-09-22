package models

type Trabalhador struct {
	ID             uint64       `gorm:"primaryKey"`
	Nome           string       `gorm:"not null;"`
	Cpf            string       `gorm:"not null; unique"`
	EmpresaID      uint64       `gorm:"not null"`                                                // Chave estrangeira que faz referência ao ID da tabela Empresa
	Empresa        Empresa      `gorm:"foreignKey:EmpresaID;constraint:OnDelete:SET NULL;"`      // Define a relação com Empresa usando EmpresaID
	DepartamentoID uint64       `gorm:"not null"`                                                // Chave estrangeira que faz referência ao ID da tabela Departamento
	Departamento   Departamento `gorm:"foreignKey:DepartamentoID;constraint:OnDelete:SET NULL;"` // Define a relação com Departamento usando DepartamentoID
}

type Empresa struct {
	ID          uint64 `gorm:"primaryKey"`
	RazaoSocial string `gorm:"not null; unique"`
	Cnpj        string `gorm:"not null;"`
}

type Departamento struct {
	ID   uint64 `gorm:"primaryKey"`
	Nome string `gorm:"not null; unique"`
}
