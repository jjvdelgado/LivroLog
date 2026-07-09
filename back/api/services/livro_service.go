package services

import (
	"api-livro/api/data/response"
	"api-livro/api/models"
	"api-livro/api/repository"
)

func ListarLivros(limit int, offset int) ([]models.Livro, error) {
	return repository.ListarLivros(limit, offset)
}

func BuscarLivro(id string) (*models.Livro, error) {
	return repository.BuscarLivroPorID(id)
}

func CriarLivro(livro *models.Livro) error {
	return repository.CriarLivro(livro)
}

func AtualizarLivro(id string, dados *models.Livro) error {
	return repository.AtualizarLivro(id, dados)
}

func RemoverLivro(id string) error {
	return repository.RemoverLivro(id)
}

func BuscarLivrosAvaliadosPorUsuario(usuarioID string) ([]models.Livro, error) {
	return repository.BuscarLivrosAvaliadosPorUsuario(usuarioID)
}

func BuscarLivroComAvaliacoes(id string) (*response.LivroComAvaliacoes, error) {
	return repository.BuscarLivroComAvaliacoes(id)
}

func PesquisarLivros(titulo, autor string, notaMin, notaMax float64) ([]response.LivroPesquisa, error) {
	return repository.PesquisarLivros(titulo, autor, notaMin, notaMax)
}

func BuscarDetalhesLivro(id string) (*response.DetalhesLivroResponse, error) {
	livro, avaliacoes, err := repository.BuscarLivroComAvaliacoesEUsuarios(id)
	if err != nil {
		return nil, err
	}

	resposta := &response.DetalhesLivroResponse{
		ID:            livro.ID.String(),
		Titulo:        livro.Titulo,
		Autor:         livro.Autor,
		AnoPublicacao: livro.AnoPublicacao,
		Sinopse:       livro.Sinopse,
		Genero:        livro.Genero,
		Editora:       livro.Editora,
		Paginas:       livro.Paginas,
		Avaliacoes:    avaliacoes,
	}

	return resposta, nil
}
