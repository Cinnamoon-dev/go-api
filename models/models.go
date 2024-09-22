package models

type Trabalhador struct {
	ID             uint64       `gorm:"primaryKey" json:"id,omitempty"`
	Nome           string       `gorm:"not null;" json:"nome,omitempty"`
	Cpf            string       `gorm:"not null; unique" json:"cpf,omitempty"`
	EmpresaID      uint64       `gorm:"not null" json:"empresa_id,omitempty"`                                                                // Chave estrangeira para Empresa
	Empresa        Empresa      `gorm:"foreignKey:EmpresaID;references:ID;constraint:OnDelete:SET NULL;" json:"empresa,omitempty"`           // Define a relação com Empresa
	DepartamentoID uint64       `gorm:"not null" json:"departamento_id,omitempty"`                                                           // Chave estrangeira para Departamento
	Departamento   Departamento `gorm:"foreignKey:DepartamentoID;references:ID;constraint:OnDelete:SET NULL;" json:"departamento,omitempty"` // Define a relação com Departamento
}

type Empresa struct {
	ID          uint64 `gorm:"primaryKey" json:"id,omitempty"`
	RazaoSocial string `gorm:"not null; unique" json:"razao_social,omitempty"`
	Cnpj        string `gorm:"not null;unique" json:"cnpj,omitempty"`
}

type Departamento struct {
	ID   uint64 `gorm:"primaryKey" json:"id,omitempty"`
	Nome string `gorm:"not null; unique" json:"nome,omitempty"`
}
