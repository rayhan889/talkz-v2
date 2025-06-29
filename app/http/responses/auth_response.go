package responses

import "github.com/google/uuid"

type (
	RegisterReponse struct {
		ID        uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt string    `json:"created_at"`
	}
)
