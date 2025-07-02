package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rayhan889/talkz-v2/app/constants"
	"github.com/rayhan889/talkz-v2/app/helpers"
	"github.com/rayhan889/talkz-v2/app/http/requests"
	"github.com/rayhan889/talkz-v2/app/models"
	"github.com/rayhan889/talkz-v2/app/repositories"
	"github.com/rayhan889/talkz-v2/app/resources"
	"github.com/rayhan889/talkz-v2/config"
	"github.com/rayhan889/talkz-v2/pkg/hash"
	"github.com/rayhan889/talkz-v2/pkg/logger"
)

type AuthService struct {
	userService            *UserService
	mailService            *MailService
	refreshTokenRepository *repositories.RefreshTokenRepository
}

func NewAuthService(
	userService *UserService,
	mailService *MailService,
	refreshTokenRepository *repositories.RefreshTokenRepository,
) *AuthService {
	return &AuthService{
		userService:            userService,
		mailService:            mailService,
		refreshTokenRepository: refreshTokenRepository,
	}
}

func (service *AuthService) Login(request *requests.LoginRequest) (string, string, error) {
	user, err := service.userService.GetByEmail(request.Email)

	userId := user.ID.String()

	if err != nil {
		return "", "", errors.New(constants.InvalidEmailOrPassword)
	}

	accessToken, err := service.GenerateAccessToken(userId)

	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := service.GenerateRefreshToken(userId)

	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken.Token, nil
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

	go func() {
		err := service.mailService.SendMail(
			user.Email,
			"Welcome to Talkz",
			resources.WelcomeEmailTemplate,
			map[string]interface{}{
				"Username": user.Username,
			},
		)
		if err != nil {
			logger.Log.Errorf("Failed to send welcome email: ", err)
			return
		}
	}()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *AuthService) RefreshToken(request *requests.RefreshTokenRequest) (string, string, error) {
	refreshToken, err := service.refreshTokenRepository.FindByToken(request.Token)

	if err != nil {
		return "", "", errors.New(constants.RefreshTokenNotFound)
	}

	if refreshToken.ValidUntil.Before(time.Now()) {
		return "", "", errors.New(constants.RefreshTokenExpired)
	}

	userId := refreshToken.UserID.String()

	accessToken, err := service.GenerateAccessToken(userId)

	if err != nil {
		return "", "", err
	}

	err = service.refreshTokenRepository.Delete(refreshToken)

	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := service.GenerateRefreshToken(userId)

	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken.Token, nil
}

func (service *AuthService) GenerateAccessToken(userId string) (string, error) {
	expTime := time.Now().Add(time.Second * time.Duration(config.Envs.JWT.Expires)).Unix()
	secretKey := []byte(config.Envs.JWT.Secret)

	claims := jwt.MapClaims{
		"sub": userId,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"exp": expTime,
		"iss": "talkz-v2",
		"aud": "talkz-v2",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *AuthService) GenerateRefreshToken(userId string) (*models.RefreshToken, error) {
	expTime := time.Now().Add(time.Second * time.Duration(config.Envs.JWT.RefreshExpires))

	userRefreshTokens, err := service.refreshTokenRepository.FindByUserId(userId)

	if err != nil {
		return nil, err
	}

	if len(userRefreshTokens) > 0 {
		for _, token := range userRefreshTokens {
			service.refreshTokenRepository.Delete(&token)
		}
	}

	newRefreshToken := models.RefreshToken{
		ID:         uuid.New(),
		UserID:     uuid.MustParse(userId),
		Token:      helpers.GenerateRandomString(32),
		ValidUntil: expTime,
	}

	err = service.refreshTokenRepository.Create(&newRefreshToken)

	if err != nil {
		return nil, err
	}

	return &newRefreshToken, err
}

func (service *AuthService) ValidateAccessToken(token string) (*models.User, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(constants.InvalidAccessTokenSigningMethod)
		}

		return []byte(config.Envs.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		userId, err := uuid.Parse(claims["sub"].(string))

		if err != nil {
			return nil, err
		}

		user, err := service.userService.GetByID(userId.String())

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, errors.New("invalid access token")
}
