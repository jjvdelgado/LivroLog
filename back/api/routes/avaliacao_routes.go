package routes

import (
	"api-livro/api/controller"

	"github.com/gin-gonic/gin"
)

func RegistrarRotasAvaliacao(r *gin.Engine) {
	r.POST("/avaliacoes", controller.CriarAvaliacao)
	r.GET("/avaliacoes", controller.ListarAvaliacoes)
	r.GET("/usuarios/:id/avaliacoes", controller.ListarAvaliacoesPorUsuario)
	r.PUT("/avaliacoes/:id", controller.AtualizarAvaliacao)
	r.DELETE("/avaliacoes/:id", controller.RemoverAvaliacao)
	r.GET("/livros/:id/avaliacoes", controller.ListarAvaliacoesComUsuarioPorLivro)

}
