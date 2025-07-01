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

func (repo *BlogRepository) FindAll() ([]models.Blog, error) {
	var blogs []models.Blog

	err := repo.db.Limit(5).Order("created_at DESC").Find(&blogs).Error

	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (repo *BlogRepository) FindBySlug(slug string) ([]models.Blog, error) {
	var blogs []models.Blog

	query := slug + "%"

	err := repo.db.Model(&models.Blog{}).Where("slug LIKE ?", query).Find(&blogs).Error

	if err != nil {
		return nil, err
	}

	return blogs, nil
}
