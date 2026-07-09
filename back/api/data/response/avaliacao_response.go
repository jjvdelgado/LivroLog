package response

import "github.com/google/uuid"

type AvaliacaoComUsuario struct {
	ID         uuid.UUID `json:"id"`
	Nota       int       `json:"nota"`
	Comentario string    `json:"comentario"`

	Usuario struct {
		ID    uuid.UUID `json:"id"`
		Nome  string    `json:"nome"`
		Email string    `json:"email"`
	} `json:"usuario"`
}
