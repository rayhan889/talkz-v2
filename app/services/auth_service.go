package services

import (
	"errors"

	"github.com/rayhan889/talkz-v2/app/http/requests"
	"github.com/rayhan889/talkz-v2/app/models"
	"github.com/rayhan889/talkz-v2/pkg/hash"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

func (service *AuthService) Register(request *requests.RegisterRequest) (*models.User, error) {
	if service.userService.IsEmailExist(request.Email) {
		return nil, errors.New("email already exists")
	}

	hash, err := hash.Make(request.Password)

	if err != nil {
		return nil, err
	}

	user, err := service.userService.CreateUser(
		request.Username,
		request.Email,
		hash,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
