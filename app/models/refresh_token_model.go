package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID         uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	Token      string    `gorm:"type:text;unique;not null"`
	ValidUntil time.Time `gorm:"type:timestamp;not null"`
	UserID     uuid.UUID
	CreatedAt  time.Time `gorm:"type:timestamp;not null;default:now()"`
}
