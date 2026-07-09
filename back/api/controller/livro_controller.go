package controller

import (
	"net/http"
	"strconv"

	"api-livro/api/models"
	"api-livro/api/services"

	"github.com/gin-gonic/gin"
)

func ListarLivros(c *gin.Context) {
	limit := 10
	offset := 0

	if l := c.Query("limit"); l != "" {
		if val, err := strconv.Atoi(l); err == nil {
			limit = val
		}
	}

	if o := c.Query("offset"); o != "" {
		if val, err := strconv.Atoi(o); err == nil {
			offset = val
		}
	}

	livros, err := services.ListarLivros(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar livros"})
		return
	}

	c.JSON(http.StatusOK, livros)
}

func BuscarLivro(c *gin.Context) {
	id := c.Param("id")
	livro, err := services.BuscarLivro(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, livro)
}

func CriarLivro(c *gin.Context) {
	var livro models.Livro
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	if err := services.CriarLivro(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Erro ao criar livro"})
		return
	}

	c.JSON(http.StatusOK, livro)
}

func AtualizarLivro(c *gin.Context) {
	id := c.Param("id")
	var dados models.Livro

	if err := c.ShouldBindJSON(&dados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	if err := services.AtualizarLivro(id, &dados); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Livro não encontrado ou erro ao atualizar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Livro atualizado com sucesso"})
}

func RemoverLivro(c *gin.Context) {
	id := c.Param("id")

	if err := services.RemoverLivro(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Livro não encontrado ou erro ao deletar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Livro removido com sucesso"})
}

func BuscarLivrosAvaliadosPorUsuario(c *gin.Context) {
	usuarioID := c.Param("id")

	livros, err := services.BuscarLivrosAvaliadosPorUsuario(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar livros avaliados"})
		return
	}

	c.JSON(http.StatusOK, livros)
}

func ListarAvaliacoesComUsuarioPorLivro(c *gin.Context) {
	livroID := c.Param("id")

	avaliacoes, err := services.BuscarAvaliacoesComUsuarioPorLivro(livroID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar avaliações"})
		return
	}

	c.JSON(http.StatusOK, avaliacoes)
}
func BuscarLivroComAvaliacoes(c *gin.Context) {
	id := c.Param("id")

	livro, err := services.BuscarLivroComAvaliacoes(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Livro não encontrado ou erro ao buscar avaliações"})
		return
	}

	c.JSON(http.StatusOK, livro)
}

func PesquisarLivros(c *gin.Context) {
	titulo := c.Query("titulo")
	autor := c.Query("autor")
	notaMinStr := c.Query("nota_min")
	notaMaxStr := c.Query("nota_max")

	var notaMin, notaMax float64
	var err error

	if notaMinStr != "" {
		notaMin, err = strconv.ParseFloat(notaMinStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": "nota_min inválida"})
			return
		}
	}

	if notaMaxStr != "" {
		notaMax, err = strconv.ParseFloat(notaMaxStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": "nota_max inválida"})
			return
		}
	}

	livros, err := services.PesquisarLivros(titulo, autor, notaMin, notaMax)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar livros"})
		return
	}

	c.JSON(http.StatusOK, livros)
}

func GetLivroDetalhes(c *gin.Context) {
	id := c.Param("id")

	livro, err := services.BuscarDetalhesLivro(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, livro)
}
