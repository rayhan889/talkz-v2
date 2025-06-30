package repositories

import (
	"github.com/rayhan889/talkz-v2/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
