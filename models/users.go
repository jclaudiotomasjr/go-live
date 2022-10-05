package models

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	Cpf  string `json:"cpf"`
	Setor string `json:"setor"`
	Pontuacao string `json:"pontuacao"`
}

