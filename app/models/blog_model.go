package models

import "github.com/google/uuid"

type Blog struct {
	ID        uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	Title     string    `gorm:"type:text;not null"`
	Slug      string    `gorm:"type:text;not null;unique"`
	Content   string    `gorm:"type:text;not null"`
	AuthorID  uuid.UUID
	CreatedAt string `gorm:"type:timestamp;not null;default:now()"`
}
