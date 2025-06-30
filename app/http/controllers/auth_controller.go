package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/constants"
	"github.com/rayhan889/talkz-v2/app/exceptions"
	"github.com/rayhan889/talkz-v2/app/helpers"
	"github.com/rayhan889/talkz-v2/app/http/requests"
	"github.com/rayhan889/talkz-v2/app/http/responses"
	"github.com/rayhan889/talkz-v2/app/services"
	"github.com/rayhan889/talkz-v2/config"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (controller *AuthController) Login(c *gin.Context) {
	loginRequest := new(requests.LoginRequest)

	err := c.BindJSON(loginRequest)
	if err != nil {
		exceptions.BadRequestError(c, errors.New(constants.ErrorInvalidRequestBody))
		return
	}

	errs := helpers.ValidateStruct(loginRequest)
	if errs != nil {
		exceptions.NewValidationError(c, errs, loginRequest)
		return
	}

	token, err := controller.authService.Login(loginRequest)

	if err != nil {
		if err.Error() == constants.InvalidEmailOrPassword {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": constants.InvalidEmailOrPassword,
				"errors":  err.Error(),
			})
			return
		}
		exceptions.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"data": responses.TokenResponse{
			AccessToken:          token,
			AccessTokenExpiresIn: config.Envs.JWT.Expires,
		},
	})
}

func (controller *AuthController) Register(c *gin.Context) {
	registerRequst := new(requests.RegisterRequest)

	err := c.BindJSON(registerRequst)
	if err != nil {
		exceptions.BadRequestError(c, errors.New(constants.ErrorInvalidRequestBody))
		return
	}

	errs := helpers.ValidateStruct(registerRequst)
	if errs != nil {
		exceptions.NewValidationError(c, errs, registerRequst)
		return
	}

	user, err := controller.authService.Register(registerRequst)

	if err != nil {
		if err.Error() == constants.ErrorEmailAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{
				"message": constants.ErrorEmailAlreadyExists,
				"errors":  err.Error(),
			})
			return
		}
		exceptions.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data": responses.RegisterReponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
		},
		"errors": nil,
	})
}

func (controller *AuthController) Logout(c *gin.Context) error {
	return nil
}
