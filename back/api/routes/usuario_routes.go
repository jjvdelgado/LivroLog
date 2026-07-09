package routes

import (
	"api-livro/api/controller"

	"github.com/gin-gonic/gin"
)

func RegistrarRotasUsuario(r *gin.Engine) {
	r.POST("/usuarios", controller.CadastrarUsuario)
	r.POST("/login", controller.Login)
	r.GET("/usuarios/:id", controller.BuscarUsuario)
	r.PUT("/usuarios/:id", controller.AtualizarUsuario)
	r.DELETE("/usuarios/:id", controller.RemoverUsuario)
	r.GET("/usuarios/:id/detalhes", controller.GetUsuarioDetalhes)
}
