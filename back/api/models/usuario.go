package models

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	Senha     string    `json:"senha"`
	Descricao string    `json:"descricao"`
	Permissao string    `json:"permissao"` // <-- NOVO CAMPO
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
