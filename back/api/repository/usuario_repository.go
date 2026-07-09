package repository

import (
	"api-livro/api/data/response"
	"api-livro/api/models"
	db "api-livro/config"
)

func CriarUsuario(usuario *models.Usuario) error {
	return db.DB.Create(usuario).Error
}

func BuscarUsuarioPorEmailESenha(email, senha string) (*models.Usuario, error) {
	var usuario models.Usuario
	result := db.DB.Where("email = ? AND senha = ?", email, senha).First(&usuario)
	if result.Error != nil {
		return nil, result.Error
	}
	return &usuario, nil
}

func BuscarUsuarioPorID(id string) (*models.Usuario, error) {
	var usuario models.Usuario
	result := db.DB.First(&usuario, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &usuario, nil
}

func AtualizarUsuario(id string, dadosAtualizados *models.Usuario) error {
	var usuario models.Usuario
	if err := db.DB.First(&usuario, "id = ?", id).Error; err != nil {
		return err
	}

	usuario.Nome = dadosAtualizados.Nome
	usuario.Email = dadosAtualizados.Email
	usuario.Senha = dadosAtualizados.Senha

	return db.DB.Save(&usuario).Error
}

func RemoverUsuario(id string) error {
	return db.DB.Delete(&models.Usuario{}, "id = ?", id).Error
}

func BuscarUsuarioComAvaliacoes(id string) (*response.UsuarioComAvaliacoes, error) {
	var usuario models.Usuario
	if err := db.DB.First(&usuario, "id = ?", id).Error; err != nil {
		return nil, err
	}

	var livros []response.LivroAvaliado

	query := `
        SELECT 
            l.id, l.titulo, l.autor, l.descricao,
            a.nota, a.comentario
        FROM livros l
        JOIN avaliacoes a ON a.livro_id = l.id
        WHERE a.usuario_id = ?
    `
	if err := db.DB.Raw(query, id).Scan(&livros).Error; err != nil {
		return nil, err
	}

	return &response.UsuarioComAvaliacoes{
		ID:              usuario.ID,
		Nome:            usuario.Nome,
		Email:           usuario.Email,
		LivrosAvaliados: livros,
	}, nil
}

func BuscarUsuarioComAvaliacoesELivros(id string) (models.Usuario, []models.Avaliacao, error) {
	var usuario models.Usuario
	if err := db.DB.First(&usuario, "id = ?", id).Error; err != nil {
		return usuario, nil, err
	}

	var avaliacoes []models.Avaliacao
	err := db.DB.
		Preload("Livro").
		Where("usuario_id = ?", id).
		Find(&avaliacoes).Error

	return usuario, avaliacoes, err
}
