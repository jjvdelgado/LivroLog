package models

import (
	"github.com/google/uuid"
)

type Livro struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Titulo        string    `json:"titulo"`
	Autor         string    `json:"autor"`
	AnoPublicacao int       `json:"ano_publicacao"`
	Sinopse       string    `json:"sinopse"`
	Genero        string    `json:"genero"`
	Editora       string    `json:"editora"`
	Paginas       int       `json:"paginas"`
}
