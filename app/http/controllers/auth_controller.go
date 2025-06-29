package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rayhan889/talkz-v2/app/helpers"
	"github.com/rayhan889/talkz-v2/app/http/requests"
	"github.com/rayhan889/talkz-v2/app/http/responses"
	"github.com/rayhan889/talkz-v2/app/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (controller *AuthController) Register(c *gin.Context) {
	registerRequst := new(requests.RegisterRequest)

	helpers.ValidateRequest(c, registerRequst)

	user, err := controller.authService.Register(registerRequst)

	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Email already exists",
				"errors":  err.Error(),
			})
			return
		}
		helpers.InternalServerError(c, err)
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

func (controller *AuthController) Login(c *gin.Context) error {
	return nil
}

func (controller *AuthController) Logout(c *gin.Context) error {
	return nil
}
