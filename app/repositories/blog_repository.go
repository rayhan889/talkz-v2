package repositories

import (
	"github.com/rayhan889/talkz-v2/app/models"
	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		db: db,
	}
}

func (repo *BlogRepository) Create(blog *models.Blog) error {
	return repo.db.Create(blog).Error
}
