package repositories

import (
	"github.com/rayhan889/talkz-v2/app/models"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (repo *RefreshTokenRepository) FindByToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken

	err := repo.db.Where("token = ?", token).First(&refreshToken).Error
	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (repo *RefreshTokenRepository) FindByUserId(userId string) ([]models.RefreshToken, error) {
	var refreshToken []models.RefreshToken

	err := repo.db.Where("user_id = ?", userId).Find(&refreshToken).Error
	if err != nil {
		return nil, err
	}

	return refreshToken, nil
}

func (repo *RefreshTokenRepository) Create(refreshToken *models.RefreshToken) error {
	return repo.db.Create(refreshToken).Error
}

func (repo *RefreshTokenRepository) Delete(refreshToken *models.RefreshToken) error {
	return repo.db.Delete(refreshToken).Error
}
