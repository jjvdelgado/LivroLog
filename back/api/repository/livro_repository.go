package repository

import (
	"api-livro/api/data/response"
	"api-livro/api/models"
	db "api-livro/config"
	"fmt"
	"strings"
)

func ListarLivros(limit int, offset int) ([]models.Livro, error) {
	var livros []models.Livro
	result := db.DB.Limit(limit).Offset(offset).Find(&livros)
	return livros, result.Error
}

func BuscarLivroPorID(id string) (*models.Livro, error) {
	var livro models.Livro
	result := db.DB.First(&livro, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &livro, nil
}

func CriarLivro(livro *models.Livro) error {
	return db.DB.Create(livro).Error
}

func AtualizarLivro(id string, dados *models.Livro) error {
	var livro models.Livro
	if err := db.DB.First(&livro, "id = ?", id).Error; err != nil {
		return err
	}

	livro.Titulo = dados.Titulo
	livro.Autor = dados.Autor
	livro.Sinopse = dados.Sinopse

	return db.DB.Save(&livro).Error
}

func RemoverLivro(id string) error {
	return db.DB.Delete(&models.Livro{}, "id = ?", id).Error
}

func BuscarLivrosAvaliadosPorUsuario(usuarioID string) ([]models.Livro, error) {
	var livros []models.Livro

	query := `
        SELECT DISTINCT l.* FROM livros l
        JOIN avaliacoes a ON a.livro_id = l.id
        WHERE a.usuario_id = ?
    `

	err := db.DB.Raw(query, usuarioID).Scan(&livros).Error
	return livros, err
}

func BuscarLivroComAvaliacoes(id string) (*response.LivroComAvaliacoes, error) {
	var livro models.Livro
	if err := db.DB.First(&livro, "id = ?", id).Error; err != nil {
		return nil, err
	}

	var avaliacoes []response.LivroAvaliacaoComUsuario

	query := `
        SELECT 
            a.id, a.nota, a.comentario,
            u.id AS "usuario.id",
            u.nome AS "usuario.nome",
            u.email AS "usuario.email"
        FROM avaliacoes a
        JOIN usuarios u ON u.id = a.usuario_id
        WHERE a.livro_id = ?
    `

	if err := db.DB.Raw(query, id).Scan(&avaliacoes).Error; err != nil {
		return nil, err
	}

	return &response.LivroComAvaliacoes{
		ID:         livro.ID,
		Titulo:     livro.Titulo,
		Autor:      livro.Autor,
		Sinopse:    livro.Sinopse,
		Avaliacoes: avaliacoes,
	}, nil
}

func PesquisarLivros(titulo, autor string, notaMin, notaMax float64) ([]response.LivroPesquisa, error) {
	var livros []response.LivroPesquisa

	var filtros []string
	if titulo != "" {
		filtros = append(filtros, fmt.Sprintf("LOWER(l.titulo) LIKE LOWER('%%%s%%')", titulo))
	}
	if autor != "" {
		filtros = append(filtros, fmt.Sprintf("LOWER(l.autor) LIKE LOWER('%%%s%%')", autor))
	}
	if notaMin > 0 {
		filtros = append(filtros, fmt.Sprintf("AVG(a.nota) >= %.2f", notaMin))
	}
	if notaMax > 0 {
		filtros = append(filtros, fmt.Sprintf("AVG(a.nota) <= %.2f", notaMax))
	}

	query := `
        SELECT l.id, l.titulo, l.autor, l.descricao, COALESCE(AVG(a.nota), 0) AS media_nota
        FROM livros l
        LEFT JOIN avaliacoes a ON a.livro_id = l.id
    `
	if len(filtros) > 0 {
		query += " WHERE " + strings.Join(filtros, " AND ")
	}
	query += " GROUP BY l.id"

	err := db.DB.Raw(query).Scan(&livros).Error
	return livros, err
}

func BuscarLivroComAvaliacoesEUsuarios(id string) (models.Livro, []models.Avaliacao, error) {
	var livro models.Livro

	if err := db.DB.First(&livro, "id = ?", id).Error; err != nil {
		return livro, nil, err
	}

	var avaliacoes []models.Avaliacao
	err := db.DB.Preload("Usuario").Where("livro_id = ?", id).Find(&avaliacoes).Error
	return livro, avaliacoes, err
}
