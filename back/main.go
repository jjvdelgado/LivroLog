package main

import (
	"api-livro/api/routes"
	db "api-livro/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Conectar()

	r := gin.Default()

	r.Use(cors.Default())

	routes.RegistrarRotasUsuario(r)
	routes.RegistrarRotasLivro(r)
	routes.RegistrarRotasAvaliacao(r)

	r.Run(":8080")
}
