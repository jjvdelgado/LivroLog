package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api-livro/api/models"
)

var DB *gorm.DB

func Conectar() {
	dsn := "host=localhost user=livro_user password=1234 dbname=livros port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}

	DB.AutoMigrate(&models.Usuario{}, &models.Livro{}, &models.Avaliacao{})
}
