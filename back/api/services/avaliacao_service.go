package services

import (
	"api-livro/api/data/response"
	"api-livro/api/models"
	"api-livro/api/repository"
)

func CriarAvaliacao(a *models.Avaliacao) error {
	return repository.CriarAvaliacao(a)
}

func ListarAvaliacoes() ([]models.Avaliacao, error) {
	return repository.ListarAvaliacoes()
}

func ListarAvaliacoesPorLivro(livroID string) ([]models.Avaliacao, error) {
	return repository.ListarAvaliacoesPorLivro(livroID)
}

func ListarAvaliacoesPorUsuario(usuarioID string) ([]models.Avaliacao, error) {
	return repository.ListarAvaliacoesPorUsuario(usuarioID)
}

func AtualizarAvaliacao(id string, dados *models.Avaliacao) error {
	return repository.AtualizarAvaliacao(id, dados)
}

func RemoverAvaliacao(id string) error {
	return repository.RemoverAvaliacao(id)
}

func BuscarAvaliacoesComUsuarioPorLivro(livroID string) ([]response.AvaliacaoComUsuario, error) {
	return repository.BuscarAvaliacoesComUsuarioPorLivro(livroID)
}
