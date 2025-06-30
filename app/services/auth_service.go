package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rayhan889/talkz-v2/app/constants"
	"github.com/rayhan889/talkz-v2/app/http/requests"
	"github.com/rayhan889/talkz-v2/app/models"
	"github.com/rayhan889/talkz-v2/config"
	"github.com/rayhan889/talkz-v2/pkg/hash"
	jwtPkg "github.com/rayhan889/talkz-v2/pkg/jwt"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

func (service *AuthService) Login(request *requests.LoginRequest) (string, error) {
	user, err := service.userService.GetByEmail(request.Email)

	if err != nil {
		return "", errors.New(constants.InvalidEmailOrPassword)
	}

	token, err := service.GenerateAccessToken(user.ID.String(), user.Username)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *AuthService) Register(request *requests.RegisterRequest) (*models.User, error) {
	if service.userService.IsEmailExist(request.Email) {
		return nil, errors.New(constants.ErrorEmailAlreadyExists)
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

func (service *AuthService) GenerateAccessToken(userId string, username string) (string, error) {
	expTime := time.Now().Add(time.Second * time.Duration(config.Envs.JWT.Expires)).Unix()
	secretKey := []byte(config.Envs.JWT.Secret)

	claims := jwtPkg.JWTClaims{
		UserID:   userId,
		Username: username,
		MapClaims: jwt.MapClaims{
			"sub": userId,
			"iat": time.Now().Unix(),
			"nbf": time.Now().Unix(),
			"exp": expTime,
			"iss": "talkz-v2",
			"aud": "talkz-v2",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
