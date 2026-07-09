package controller

import (
	"api-livro/api/models"
	"api-livro/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CadastrarUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	usuario.Permissao = "usuario"

	if err := services.CadastrarUsuario(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Erro ao cadastrar"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func Login(c *gin.Context) {
	var login struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	usuario, err := services.LoginUsuario(login.Email, login.Senha)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"erro": "Login inválido"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func BuscarUsuario(c *gin.Context) {
	id := c.Param("id")

	usuario, err := services.BuscarUsuario(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func AtualizarUsuario(c *gin.Context) {
	id := c.Param("id")
	var dados models.Usuario

	if err := c.ShouldBindJSON(&dados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	dados.Permissao = "usuario"

	if err := services.AtualizarUsuario(id, &dados); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado ou erro ao atualizar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário atualizado com sucesso"})
}

func RemoverUsuario(c *gin.Context) {
	id := c.Param("id")

	if err := services.RemoverUsuario(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado ou erro ao deletar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": "Usuário removido com sucesso"})
}

func BuscarUsuarioComAvaliacoes(c *gin.Context) {
	id := c.Param("id")

	usuario, err := services.BuscarUsuarioComAvaliacoes(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado ou erro ao buscar avaliações"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func GetUsuarioDetalhes(c *gin.Context) {
	id := c.Param("id")

	dados, err := services.BuscarDetalhesUsuario(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, dados)
}
