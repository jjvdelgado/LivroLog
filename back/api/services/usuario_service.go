package services

import (
	"api-livro/api/data/response"
	"api-livro/api/models"
	"api-livro/api/repository"
)

func CadastrarUsuario(usuario *models.Usuario) error {
	return repository.CriarUsuario(usuario)
}

func LoginUsuario(email, senha string) (*models.Usuario, error) {
	return repository.BuscarUsuarioPorEmailESenha(email, senha)
}

func BuscarUsuario(id string) (*models.Usuario, error) {
	return repository.BuscarUsuarioPorID(id)
}
func AtualizarUsuario(id string, dados *models.Usuario) error {
	return repository.AtualizarUsuario(id, dados)
}

func RemoverUsuario(id string) error {
	return repository.RemoverUsuario(id)
}

func BuscarUsuarioComAvaliacoes(id string) (*response.UsuarioComAvaliacoes, error) {
	return repository.BuscarUsuarioComAvaliacoes(id)
}

func BuscarDetalhesUsuario(id string) (*response.UsuarioDetalhadoResponse, error) {
	usuario, avaliacoes, err := repository.BuscarUsuarioComAvaliacoesELivros(id)
	if err != nil {
		return nil, err
	}

	// Mapa para garantir livros únicos
	livrosMap := make(map[string]models.Livro)
	for _, avaliacao := range avaliacoes {
		livroID := avaliacao.Livro.ID.String()
		if _, existe := livrosMap[livroID]; !existe {
			livrosMap[livroID] = avaliacao.Livro
		}
	}

	// Converter mapa para slice
	livros := make([]models.Livro, 0, len(livrosMap))
	for _, livro := range livrosMap {
		livros = append(livros, livro)
	}

	return &response.UsuarioDetalhadoResponse{
		ID:         usuario.ID.String(),
		Nome:       usuario.Nome,
		Email:      usuario.Email,
		Descricao:  usuario.Descricao,
		CreatedAt:  usuario.CreatedAt.Format("2006-01-02 15:04:05"),
		Avaliacoes: avaliacoes,
		Livros:     livros,
	}, nil
}
