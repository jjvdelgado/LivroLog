package repository

import (
	"api-livro/api/data/response"
	"api-livro/api/models"
	db "api-livro/config"
)

func CriarAvaliacao(a *models.Avaliacao) error {
	return db.DB.Create(a).Error
}

func ListarAvaliacoes() ([]models.Avaliacao, error) {
	var avaliacoes []models.Avaliacao
	err := db.DB.Find(&avaliacoes).Error
	return avaliacoes, err
}

func ListarAvaliacoesPorLivro(livroID string) ([]models.Avaliacao, error) {
	var avaliacoes []models.Avaliacao
	err := db.DB.Where("livro_id = ?", livroID).Find(&avaliacoes).Error
	return avaliacoes, err
}

func ListarAvaliacoesPorUsuario(usuarioID string) ([]models.Avaliacao, error) {
	var avaliacoes []models.Avaliacao
	err := db.DB.Where("usuario_id = ?", usuarioID).Find(&avaliacoes).Error
	return avaliacoes, err
}

func AtualizarAvaliacao(id string, dados *models.Avaliacao) error {
	var avaliacao models.Avaliacao
	if err := db.DB.First(&avaliacao, "id = ?", id).Error; err != nil {
		return err
	}

	avaliacao.Nota = dados.Nota
	avaliacao.Comentario = dados.Comentario

	return db.DB.Save(&avaliacao).Error
}

func RemoverAvaliacao(id string) error {
	return db.DB.Delete(&models.Avaliacao{}, "id = ?", id).Error
}

func BuscarAvaliacoesComUsuarioPorLivro(livroID string) ([]response.AvaliacaoComUsuario, error) {
	var resultado []response.AvaliacaoComUsuario

	query := `
        SELECT 
            a.id, a.nota, a.comentario,
            u.id as "usuario.id",
            u.nome as "usuario.nome",
            u.email as "usuario.email"
        FROM avaliacoes a
        JOIN usuarios u ON u.id = a.usuario_id
        WHERE a.livro_id = ?
    `

	err := db.DB.Raw(query, livroID).Scan(&resultado).Error
	return resultado, err
}
