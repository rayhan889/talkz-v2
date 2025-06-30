package services

import (
	"github.com/rayhan889/talkz-v2/app/models"
	"github.com/rayhan889/talkz-v2/app/repositories"
	"github.com/rayhan889/talkz-v2/pkg/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) CreateUser(
	username string,
	email string,
	password string) (*models.User, error) {
	user := models.User{
		ID:       uuid.GenerateUUID(),
		Username: username,
		Email:    email,
		Password: password,
	}

	err := service.userRepository.Create(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (service *UserService) GetByEmail(email string) (*models.User, error) {
	return service.userRepository.FindByEmail(email)
}

func (service *UserService) IsEmailExist(email string) bool {
	_, err := service.userRepository.FindByEmail(email)

	return err != gorm.ErrRecordNotFound
}
