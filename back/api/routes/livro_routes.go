package routes

import (
	"api-livro/api/controller"

	"github.com/gin-gonic/gin"
)

func RegistrarRotasLivro(r *gin.Engine) {
	r.GET("/livros", controller.ListarLivros)
	r.GET("/livros/:id", controller.BuscarLivro)
	r.POST("/livros", controller.CriarLivro)
	r.PUT("/livros/:id", controller.AtualizarLivro)
	r.DELETE("/livros/:id", controller.RemoverLivro)
	r.GET("/usuarios/:id/livros-avaliados", controller.BuscarLivrosAvaliadosPorUsuario)
	r.GET("/livros/pesquisa", controller.PesquisarLivros)
	r.GET("/livros/:id/detalhes", controller.GetLivroDetalhes)
}
