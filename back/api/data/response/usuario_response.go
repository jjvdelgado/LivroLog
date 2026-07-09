package response

import (
	"api-livro/api/models"

	"github.com/google/uuid"
)

type LivroAvaliado struct {
	ID         uuid.UUID `json:"id"`
	Titulo     string    `json:"titulo"`
	Autor      string    `json:"autor"`
	Descricao  string    `json:"descricao"`
	Nota       int       `json:"nota"`
	Comentario string    `json:"comentario"`
}

type UsuarioComAvaliacoes struct {
	ID              uuid.UUID       `json:"id"`
	Nome            string          `json:"nome"`
	Email           string          `json:"email"`
	LivrosAvaliados []LivroAvaliado `json:"livros_avaliados"`
}

type UsuarioDetalhadoResponse struct {
	ID         string             `json:"id"`
	Nome       string             `json:"nome"`
	Email      string             `json:"email"`
	Descricao  string             `json:"descricao"`
	CreatedAt  string             `json:"created_at"`
	Avaliacoes []models.Avaliacao `json:"avaliacoes"`
	Livros     []models.Livro     `json:"livros"`
}
