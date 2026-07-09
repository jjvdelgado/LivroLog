package models

import (
	"time"

	"github.com/google/uuid"
)

type Avaliacao struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UsuarioID  uuid.UUID `json:"usuario_id"`
	LivroID    uuid.UUID `json:"livro_id"`
	Nota       int       `json:"nota"`
	Comentario string    `json:"comentario"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`

	Usuario Usuario `gorm:"foreignKey:UsuarioID" json:"usuario"`
	Livro   Livro   `gorm:"foreignKey:LivroID" json:"livro"`
}
