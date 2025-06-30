package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	Username  string    `gorm:"type:text;not null"`
	Email     string    `gorm:"type:text;unique;not null"`
	Password  string    `gorm:"type:text;not null"`
	Blogs     []Blog    `gorm:"foreignKey:ID;references:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:now()"`
}
