package controller

import (
	"api-livro/api/models"
	"api-livro/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CriarAvaliacao(c *gin.Context) {
	var a models.Avaliacao
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	if err := services.CriarAvaliacao(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Erro ao criar avaliação"})
		return
	}

	c.JSON(http.StatusOK, a)
}

func ListarAvaliacoes(c *gin.Context) {
	avaliacoes, err := services.ListarAvaliacoes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar avaliações"})
		return
	}

	c.JSON(http.StatusOK, avaliacoes)
}

func ListarAvaliacoesPorLivro(c *gin.Context) {
	livroID := c.Param("id")
	avaliacoes, err := services.ListarAvaliacoesPorLivro(livroID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar avaliações do livro"})
		return
	}

	c.JSON(http.StatusOK, avaliacoes)
}

func ListarAvaliacoesPorUsuario(c *gin.Context) {
	usuarioID := c.Param("id")
	avaliacoes, err := services.ListarAvaliacoesPorUsuario(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar avaliações do usuário"})
		return
	}

	c.JSON(http.StatusOK, avaliacoes)
}

func AtualizarAvaliacao(c *gin.Context) {
	id := c.Param("id")
	var dados models.Avaliacao

	if err := c.ShouldBindJSON(&dados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	if err := services.AtualizarAvaliacao(id, &dados); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Erro ao atualizar avaliação"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Avaliação atualizada com sucesso"})
}

func RemoverAvaliacao(c *gin.Context) {
	id := c.Param("id")

	if err := services.RemoverAvaliacao(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Erro ao remover avaliação"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Avaliação removida com sucesso"})
}
