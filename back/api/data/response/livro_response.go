package response

import (
	"api-livro/api/models"

	"github.com/google/uuid"
)

type UsuarioAvaliador struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Email string    `json:"email"`
}

type LivroAvaliacaoComUsuario struct {
	ID         uuid.UUID        `json:"id"`
	Nota       int              `json:"nota"`
	Comentario string           `json:"comentario"`
	Usuario    UsuarioAvaliador `json:"usuario"`
}

type LivroComAvaliacoes struct {
	ID         uuid.UUID                  `json:"id"`
	Titulo     string                     `json:"titulo"`
	Autor      string                     `json:"autor"`
	Sinopse    string                     `json:"descricao"`
	Avaliacoes []LivroAvaliacaoComUsuario `json:"avaliacoes"`
}

type LivroPesquisa struct {
	ID        uuid.UUID `json:"id"`
	Titulo    string    `json:"titulo"`
	Autor     string    `json:"autor"`
	Sinopse   string    `json:"descricao"`
	MediaNota float64   `json:"media_nota"`
}

type DetalhesLivroResponse struct {
	ID            string             `json:"id"`
	Titulo        string             `json:"titulo"`
	Autor         string             `json:"autor"`
	AnoPublicacao int                `json:"ano_publicacao"`
	Sinopse       string             `json:"sinopse"`
	Genero        string             `json:"genero"`
	Editora       string             `json:"editora"`
	Paginas       int                `json:"paginas"`
	Avaliacoes    []models.Avaliacao `json:"avaliacoes"`
}
