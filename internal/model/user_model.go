package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username  string    `gorm:"type:text;not null" json:"username"`
	Email     string    `gorm:"type:text;unique;not null" json:"email"`
	Password  string    `gorm:"type:text;not null" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:now()" json:"created_at"`
}

type RegisterPayload struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
